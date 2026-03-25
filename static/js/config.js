const config = {
    init(){

        document.getElementById("form-config").addEventListener("submit", (e) => {
            e.preventDefault();
            config.update(e.target);

        });

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


    }

};

(function () {
    'use strict';
    config.init();
})();


