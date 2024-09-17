# Simple Storage Organization Management System #
- A Go+Vue+Postgres app for keeping your storage system neat, indexed, and easily searchable

## The Lore ##
- Going between houses and apartments, I find myself reorganizing my workshop and office each time I move. The ways I organize everything also isn't very portable, so after each move, a complete re-organization is required.
- This project aims to provide some time of portable lab that can easily go through moves, while still having very simple day-to-day usage.
- I wanted something that's **easily expandable** and **easily portable**. For this I chose the Ridgid Pro Gear Toolbox System. Any hardware brand's modular toolbox system can be used for this. For those looking to "ball out", I recommend the Milwaulkee PackOut system. While not required, I own a 3D printer and wanted something that had a decent amount of 3D designs available online for breaking down these toolboxes into even smaller components. Luckily, a few 3D models exist for the Ridgid system, but not as many as some of the more popular systems (like PackOut).
- I wanted a system where **finding ANYTHING is easy**. When the number of toolboxes goes up, unless you're viciously relabeling boxes or have a really good memory, finding anything (especially tools you don't use often) becomes a nightmare. That's where this Github repo comes in. `SWAGSOMS` serves as a basic software layer that can be used to easily index and search for items in your toolbox system.
    - By looking at your toolboxes as a tree, each toolbox as a parent node, and each slot in the toolbox as a child node of the parent, you can begin to see how this works. Each toolbox slot can then be broken up into smaller spots (as would be the case when 3D printing subdividers), which would become children of the slot node.
    - The Go/Gin backend uses a PostgreSQL database for storage and provides a very minimal Vue UI for searching, adding, editting, and deleting slots. This app can basically be hosted anywhere, and you can use the UI on your phone's browser or on a dedicated machine to figure out where anything in your lab is, in seconds. 
