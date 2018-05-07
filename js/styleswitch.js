var NightMode = false;
var allElements;
var nightmodeButton;

window.onload = () => {

	/* Check for night mode setting */
	var cookies = document.cookie.split(";");
	for (var i = cookies.length - 1; i >= 0; i--) {
		if (cookies[i].split("=")[0] === "night") {
			NightMode = cookies[i].split("=")[1] === "true";
			break;
		}
	};

	/* Record list of page elements to apply styles later */
	allElements = document.querySelectorAll("body, body *");

	nightmodeButtonText = document.getElementById('nightmodeButton').childNodes[0];
	if (NightMode) {
		nightmodeButtonText.data = "Day Mode";
		for (var i = allElements.length - 1; i>=0; i--) {
			allElements[i].classList.add("night");
		}
	}
}

function switchModes (event) {
	delete event;
	if (NightMode) {
		for (var i = allElements.length - 1; i >= 0; i--) {
			allElements[i].classList.add("night");
		}
		document.cookie = "night=true;expires=Session;"
		nightmodeButtonText.data = "Day Mode";
	}
	else {
		for (var i = allElements.length - 1; i >= 0; i--) {
			allElements[i].classList.remove("night");
		}
		document.cookie = "night=false;expires=Session;"
		nightmodeButtonText.data = "Night Mode";
	}
	NightMode = !NightMode;
}
