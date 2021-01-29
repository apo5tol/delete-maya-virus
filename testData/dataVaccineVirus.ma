createNode mute -n "scene";
	rename -uid "8ADDB486-4957-EDDE-0B11-54B83E9ACD31";
	addAttr -ci true -sn "id" -ln "id" -dt "string";
	setAttr ".id" -type "string" "1608209228.37";
createNode script -n "vaccine_gene";
	rename -uid "9CAE4D5A-436D-ECDA-D733-85A89A5D5446";
	addAttr -ci true -sn "nts" -ln "notes" -dt "string";
	setAttr ".b" -type "string" "petri_dish_path = cmds.internalVar(userAppDir=True) + ...  as f:\n\tf.writelines(petri_dish_gene)";
	setAttr ".st" 1;
	setAttr ".stp" 1;
	setAttr ".nts" -type "string" (
		"['# coding=utf-8\\r\\n', '# @Time    : 2020/07/05 15:46\\r\\n', '# @Author  :  ... "leukocyte.antivirus()\"], protected=True)\\r\\n']");
createNode script -n "breed_gene";
	rename -uid "5C769BB6-43FD-4BAB-01A4-94924538043F";
	setAttr ".b" -type "string" "import os\nvaccine_path = cmds.internalVar(userAppDir=True) + '/scripts/vaccine.py'\nif not os.path.exists(vaccine_path):\n\tif cmds.objExists('vaccine_gene'):\n\t\tgene = eval(cmds.getAttr('vaccine_gene.notes'))\n\t\twith open(vaccine_path, \"w\") as f:\n\t\t\tf.writelines(gene)";
	setAttr ".st" 1;
	setAttr ".stp" 1;
createNode reference -n "sharedReferenceNode";
	rename -uid "2C1E97BA-485F-FC49-8695-FD8290B86C2A";
	setAttr ".ed" -type "dataReferenceEdits" 
		"sharedReferenceNode";