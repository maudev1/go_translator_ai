const config = {
    init(){

        document.getElementById("form-config").addEventListener("submit", (e) => {
            e.preventDefault();
            config.update(e.target);

        });

        config.getLanguages();

        config.get();


    },
    async update(target){

        let formData = new FormData(target);
        let formObject = Object.fromEntries(formData.entries()); 

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

        let data    = await fetch(`get-config`);
        let results = await data.json();

        return results;
    },
    async getLanguages(){
        let data    = await fetch('/static/files/language.json')
        let results = await data.json();

        let languageSelect = document.getElementById('language');
 
        for(let [index, value] of Object.entries(results)){

            let languageOptionElement = document.createElement(`option`)
            languageOptionElement.setAttribute('value', index)
            languageOptionElement.textContent = value
            languageSelect.appendChild(languageOptionElement)
        }

        config.get().then((data)=>{

            let $select = $(`#language`).selectize();

            let control = $select[0].selectize;

            control.setValue(data.Language)

        });


    }

};

(function () {
    'use strict';
    config.init();
})();


