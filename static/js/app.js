const app = {

    currentPosition: 0,
    lastPosition: 0,

    init() {

        app.setPosition();

        app.updatePreview(app.currentPosition);

        document.getElementById('content').addEventListener('wheel', (event) => {

            app.navigate(event.deltaY);

        });

        document.querySelector('.translate-google').addEventListener('click', async function (event) {

            const engine = event.target.dataset.engine

            app.autoTranslate(engine);

        });

        document.querySelector('.translate-ai').addEventListener('click', async function (event) {

            const engine = event.target.dataset.engine

            app.autoTranslate(engine);

        });

        document.querySelector('#search-index').addEventListener('click', () => {

            let index = document.getElementById('index').value;

            if (!index || index < 1 || typeof index == "undefined") {
                return;
            }

            app.currentPosition = index;

            app.updatePreview(index);

            app.showPosition();

        });


        // Keyboard control 


        document.addEventListener('keypress', (event) => {

            console.log(event)

            let key = event.key.toUpperCase();

            switch (key) {

                case 'W': {

                    app.navigate(1);

                    return;
                }
                case 'A': {

                    app.autoTranslate('google');

                    return;
                }
                case 'S': {

                    app.navigate(2);

                    return;
                }
                case 'D': {

                    app.autoTranslate('ai');

                    return;

                }
                // case 'ENTER':{

                // }
                default: {
                    return
                }


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
    navigate(direction) {

        let downSelector = document.querySelector('.down');
        let upSelector = document.querySelector('.up');

        downSelector.classList.remove('active');
        upSelector.classList.remove('active');

        if (direction > 1) {

            // down

            downSelector.classList.add('active');

            app.setPosition('down');

        } else {

            // up 

            upSelector.classList.add('active');

            app.setPosition('up');

        }


        setTimeout(function () {
            downSelector.classList.remove('active')
            upSelector.classList.remove('active')
        }, 1000)


    },
    updatePreview(index) {

        app.wordList().then((text) => {

            document.getElementById('editor').value = text[index].value

        })


    },
    setPosition(direction) {

        app.wordList().then((texts) => {

            let initial = 0;
            let lastPosition = texts.length;
            app.lastPosition = lastPosition;

            let position = Number(app.currentPosition) + 1

            if (direction == "up") {

                position = Number(app.currentPosition) - 1
            }


            if (position <= 0 || position > lastPosition) {
                return;
            }

            if (position == lastPosition) {

                app.updatePreview(initial)

                app.currentPosition = initial;

            } else {

                app.updatePreview(position);

                app.currentPosition = position;

            }

            app.showPosition();

        })


    },
    showPosition() {
        document.getElementById('index-position').innerHTML = `${app.currentPosition}/${app.lastPosition}`
    },
    async autoTranslate(engine) {

        const text = document.getElementById('editor').value;

        let data = await fetch('/translate', {
            method: 'POST',
            body: JSON.stringify({
                engine: engine,
                text: text
            })
        })

        if (data.ok) {
            let response = await data.json();

            document.getElementById('editor').value = response.translated

        }

    },



};


(function () {

    'use strict';
    app.init();

})();


