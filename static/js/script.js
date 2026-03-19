const app = {

    currentPosition: 0,
    lastPosition: 0,

    init() {

        app.showPosition();

        app.updatePreview(app.currentPosition);

        document.getElementById('content').addEventListener('wheel', (event) => {

            if (event.deltaY > 1) {
                console.log('baixo')
            } else {
                console.log('cima')

            }

            app.selectizeWord(event.deltaY);

            app.showPosition();

        })





    },
    async wordList() {

        let data = await fetch('/load-base-text')
        let results = await data.json();

        if (results) {
            return results;
        }


        return ['uva', 'pera', 'maca', 'salada mista'];
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
            document.getElementById('editor').innerHTML = text[index].value

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


