package main;

import "database/sql";
import "fmt";
import "io/ioutil";
import "log";
import "net/http";
import "path/filepath";
import "strings";

import _ "github.com/lib/pq";
import getopt "github.com/pborman/getopt/v2";

const genericInternalError = "<html><body><h1>500 - Internal Server Error</h1><pre>%s</pre></body></html>";

// var pages = map[string]http.HandlerFunc{};
var mimes = map[string]string{};
var handlers = map[string]func(http.ResponseWriter, *http.Request){};

var db *sql.DB;

type alignment struct {
	structure string
	nature string
}

type Diety struct {
	name string
	alignment []uint8
	title string
	pantheon string
	symbol string
	source string
	reprinted bool
	page int
	description *string
};

const dietyFmt = "<tr><td>%s</td><td>%s</td><td>%s</td><td>%s</td><td>%s</td></tr>\n";

func serveFile(w http.ResponseWriter, f string) {
	if f == "/" {
		f = "5etools.html";
	}

	if strings.HasPrefix(f, "/") {
		f = f[1:];
	}

	ext := filepath.Ext(f);
	if mime, ok := mimes[ext]; ok {
		w.Header().Set("Content-Type", mime);
	}

	content, err := ioutil.ReadFile("./" + f);
	if err != nil {
		w.WriteHeader(http.StatusNotFound);
		log.Printf("File serve error: %s\n", err.Error());
		return
	}

	w.Write(content);
}

func ServeHTTP(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close();
	log.Printf("Serving request '%s %s'", req.Method, req.URL)
	if req.Method == "BREW" {
		w.WriteHeader(http.StatusTeapot);
		return;
	}
	if handler, ok := handlers[req.URL.Path]; ok {
		handler(w, req);
	} else if req.Method == "GET" {
		log.Printf("No special handler for '%s' - serving as file\n", req.URL.Path);
		serveFile(w, req.URL.Path);
	} else {
		w.WriteHeader(http.StatusNotFound);
	}
}

func dieties(w http.ResponseWriter, req *http.Request) {
	rows, err := db.Query("SELECT * FROM diety;");
	if err != nil {
		w.WriteHeader(http.StatusBadGateway);
		w.Write([]byte("Failed to connect to database, unknown error\n"));
		fmt.Printf("Database retrieval error: %s\n", err.Error());
		return
	}

	c, err := ioutil.ReadFile("./deities.html");
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError);
		w.Write([]byte(fmt.Sprintf(genericInternalError, err.Error())));
		log.Printf("Error reading deities.html: %s\n", err.Error());
		return;
	}

	content := string(c);
	i := strings.Index(content, "{{deities}}");
	if i < 0 {
		log.Printf("Error reading deities.html: template string not found!\n");
		log.Printf("Contents:\n=====================\n%s\n======================\n", string(content));
		w.WriteHeader(http.StatusInternalServerError);
		w.Write([]byte(fmt.Sprintf(genericInternalError, "template string not found!\n")));
		return;
	}

	w.Write([]byte(content[:i]));
	content = content[i+11:];

	for rows.Next() {
		row := Diety{};
		err := rows.Scan(&row.name, &row.alignment, &row.title, &row.pantheon, &row.symbol, &row.source, &row.reprinted, &row.page, &row.description);
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError);
			w.Write([]byte(fmt.Sprintf(genericInternalError, err.Error())));
			log.Printf("Error reading database row: %s\n", err.Error());
			return;
		}
		w.Write([]byte(fmt.Sprintf(dietyFmt, row.name, row.pantheon, row.alignment, row.title, row.source)));
	}

	w.Write([]byte(content));

}

// Command line flags


func main() {
	dbName := getopt.StringLong("dbname", 'd', "dnd", "Database name to connect to");
	dbHost := getopt.StringLong("host", 'h', "localhost", "Database server host");
	dbUser := getopt.StringLong("username", 'U', "postgres", "Database user name");
	dbPort := getopt.UintLong("port", 'p', 5432, "Database server port");
	dbPasswd := getopt.StringLong("password", 'W', "", "Database password");

	getopt.Parse();
	mimes [".html"] = "text/html";
	mimes [".css"] = "text/css";
	mimes [".js"] = "text/javascript";

	psqlInfo := fmt.Sprintf("host=%s user=%s password=%s port=%d dbname=%s sslmode=require", *dbHost, *dbUser, *dbPasswd, *dbPort, *dbName);
	log.Printf("Connecting to database with: %s", psqlInfo);
	var err error;
	if db, err = sql.Open("postgres", psqlInfo); err != nil {
		log.Fatalf("%s\n", err.Error());
	}
	if err = db.Ping(); err != nil {
		log.Fatalf("%s\n", err.Error());
	}

	defer db.Close();

	handlers["/deities.html"] = dieties;
	s := &http.Server {
		Addr: ":8080",
		Handler: http.HandlerFunc(ServeHTTP),
	};
	if err := s.ListenAndServe(); err != nil {
		log.Fatalf("%s\n", err.Error());
	}
}
