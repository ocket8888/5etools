## 5e Tools
I'm trying to overhaul 5etools.com to not use node.js
*node.js bad*

## Running 5etools Locally (Offline Copy)
There are several options for running a local/offline copy of 5etools, including:

**Beginner:** Use Firefox to open the files.

**Advanced:** Host the project locally on a dev webserver, perhaps using [this](https://github.com/cortesi/devd).

## How to import 5etools beasts/spells/items into Roll20
1. Get Greasemonkey (Firefox) or Tampermonkey (Chrome).

2. Click [here](https://github.com/TheGiddyLimit/5etoolsR20/raw/master/5etoolsR20.user.js) and install the script.

3. Open the Roll20 game where you want the stuff imported.

4. Go to the gear icon and click on the things you want imported.

5. Let it run. The journal will start fill up with the stuff you selected. It's not too laggy but can take a long time depending on the amount of stuff you selected.

6. Bam. Done. If you are using the Shaped sheet, be sure to open up the NPC sheets and let them convert before using it.

You can convert stat blocks to JSON for importing via [this converter](converter.html).

## Dev Notes

### Style Guidelines
- Use tabs over spaces.

### JSON Cleaning
#### Trailing Commas
Don't leave commas at the end of objects/arrays - they don't serve any purpose and merely make the file bigger than it needs to be.

#### Character Set
- Use UTF-8 encoding
- ’ should be replaced with '
- “ and ” should be replaced with "
- — (em dash) should be replaced with \u2014 (Unicode for em dash)
- – and \u2013 (en dash) should be replaced with \u2014
- the only Unicode escape sequence allowed is \u2014; all other characters (unless noted above) should be stored as-is

#### Convention for dashes
- - (hyphen) should **only** be used to hyphenate words, e.g. 60-foot and 18th-level
- any whitespace on any side of a \u2014 should be removed

#### Convention for measurement
- Adjectives: a hyphen and the full name of the unit of measure should be used, e.g. dragon exhales acid in a 60-foot line
- Nouns: a space and the short name of the unit of measure (including the trailing period) should be used, e.g. blindsight 60 ft., darkvision 120 ft.
- Time: a slash, /, with no spaces on either side followed by the capitalised unit of time, e.g. 2/Turn, 3/Day

## License

This project is licensed under the terms of the MIT license.
