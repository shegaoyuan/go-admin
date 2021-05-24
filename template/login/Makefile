all:
	find ./ -name ".DS_Store" -depth -exec rm {} \;
	adm combine js -s ./assets/src/js/combine/ -d ./assets/login/dist/all.min.js
	adm combine css -s ./assets/src/css/ -d ./assets/login/dist/all.min.css
	adm compile asset -s ./assets/login/dist/ -d ./ -p=login