const config = {
    init() {

        document.getElementById("form-config").addEventListener("submit", (e) => {
            e.preventDefault();
            config.update(e.target);

        });

        document.getElementById("file-uploader").addEventListener("submit", (e) => {
            e.preventDefault();
            config.fileUploader(e.target);

        });

        config.getLanguages();

        config.get();


    },
    async update(target) {

        let formData = new FormData(target);
        let formObject = Object.fromEntries(formData.entries());

        let options = {
            method: "POST",
            body: JSON.stringify(formObject, null, 2),
            headers: {
                'Content-Type': 'application/json',
            }
        }

        let data = await fetch('set-config', options)

        if (data) {

            helpers.custom_alert("The Configuration has been saved", "success");
        }

    },

    async fileUploader(target) {

        const formData = new FormData();

        formData.append("baseFile", target[0].files[0]);

        let options = {
            method: "POST",
            body: formData
        }

        let data = await fetch('set-basefile', options)

    },
    async get() {

        let data = await fetch(`get-config`);
        let results = await data.json();

        // $(`#language`).selectize({
        //     items:[results.Language]
        // })


    },
    async getLanguages() {
        let data = await fetch('/static/files/language.json')
        let results = await data.json();

        let languageSelect = document.getElementById('language');

        for (let [index, value] of Object.entries(results)) {

            let languageOptionElement = document.createElement(`option`)
            languageOptionElement.setAttribute('value', index)
            languageOptionElement.textContent = value
            languageSelect.appendChild(languageOptionElement)
        }

        $(`#language`).selectize();

    }

};

(function () {
    'use strict';
    config.init();
})();


