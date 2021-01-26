createNode mute -n "scene";
	rename -uid "8ADDB486-4957-EDDE-0B11-54B83E9ACD31";
	addAttr -ci true -sn "id" -ln "id" -dt "string";
	setAttr ".id" -type "string" "1608209228.37";
createNode reference -n "sharedReferenceNode";
	rename -uid "2C1E97BA-485F-FC49-8695-FD8290B86C2A";
	setAttr ".ed" -type "dataReferenceEdits" 
		"sharedReferenceNode";