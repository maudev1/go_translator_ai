const app = {

    currentPosition: 0,
    lastPosition: 0,

    init() {

        app.showPosition();

        app.updatePreview(app.currentPosition);

        document.getElementById('content').addEventListener('wheel', (event) => {

            app.selectizeWord(event.deltaY);

            app.showPosition();

        });

        document.querySelector('.translate').addEventListener('click', async function(event){

            const engine = event.target.dataset.engine
            const text   = document.getElementById('editor').value;

            let data = await fetch('/translate', {
                method:'POST',
                body: JSON.stringify({
                    engine: engine,
                    text: text
                })
            }) 

            if(data.ok){
                let response = await data.json();

                document.getElementById('editor').value = response.translated

            }

  

        });


    },
    async wordList() {

        let data = await fetch('/load-base-text')
        let results = await data.json();

        if (results) {
            return results;
        }

    },
    selectizeWord(position) {

        let downSelector = document.querySelector('.down')
        let upSelector = document.querySelector('.up')

        downSelector.classList.remove('active')
        upSelector.classList.remove('active')

        if (position > 1) {

            downSelector.classList.add('active')

        } else {

            upSelector.classList.add('active')

        }

        setTimeout(function () {
            downSelector.classList.remove('active')
            upSelector.classList.remove('active')
        }, 1000)


    },
    updatePreview(index) {

        app.wordList().then((text) => {

            console.log(index)

            console.log(text[index].value)

            document.getElementById('editor').value = text[index].value

        })


    },
    showPosition() {

        app.wordList().then((texts) => {

            let initial = 0;
            let lastPosition = texts.length;
            let position = app.currentPosition + 1


            if (position == lastPosition) {

                app.updatePreview(initial)

                app.currentPosition = initial;

            } else {

                app.updatePreview(position);

                app.currentPosition = position;

            }


            document.getElementById('index-position').innerHTML = `${app.currentPosition + 1}/${lastPosition}`

        })


    }




};


(function () {

    'use strict';
    app.init();

})();


