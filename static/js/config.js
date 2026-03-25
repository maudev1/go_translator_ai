const config = {
    init(){

        document.getElementById("form-config").addEventListener("submit", (e) => {
            e.preventDefault();
            config.update(e.target);

        });

        config.getLanguages()

    },
    async update(target){

        let formData = new FormData(target);
        let formObject = Object.fromEntries(formData.entries()); 

        console.log(formObject)

        let options = {
            method:"POST",
            body:JSON.stringify(formObject, null, 2),
            headers:{
                'Content-Type': 'application/json',
            }
        }

        let data = await fetch('set-config', options)



    },
    async get(){


    },
    async getLanguages(){
        let data = await fetch('/static/files/language.json')
        let results = await data.json();

    }

};

(function () {
    'use strict';
    config.init();
})();


