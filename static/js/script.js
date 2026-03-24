const app = {

    currentPosition: 0,
    lastPosition: 0,

    init() {

        app.showPosition();

        app.updatePreview(app.currentPosition);

        document.getElementById('content').addEventListener('wheel', (event) => {

            app.navigate(event.deltaY);

            // app.showPosition();

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
    navigate(direction) {

        let downSelector = document.querySelector('.down');
        let upSelector = document.querySelector('.up');

        downSelector.classList.remove('active');
        upSelector.classList.remove('active');

        if (direction > 1) {

            // down

            downSelector.classList.add('active');

             app.showPosition('down');

        } else {

            // up 

            upSelector.classList.add('active');

            app.showPosition('up');
            
        }
        
        
        setTimeout(function () {
            downSelector.classList.remove('active')
            upSelector.classList.remove('active')
        }, 1000)


    },
    updatePreview(index) {

        // debugger/

        app.wordList().then((text) => {

            document.getElementById('editor').value = text[index].value

        })


    },
    showPosition(direction) {

        app.wordList().then((texts) => {

            let initial = 0;
            let lastPosition = texts.length;
            
            let position = app.currentPosition + 1
            
            if(direction == "down"){
                
                position = app.currentPosition - 1
            }


            if(position <= 0 || position > lastPosition){
                return;
            }

            if (position == lastPosition) {

                app.updatePreview(initial)

                app.currentPosition = initial;

            } else {

                app.updatePreview(position);

                app.currentPosition = position;

            }


            document.getElementById('index-position').innerHTML = `${position}/${lastPosition}`

        })


    }





};


(function () {

    'use strict';
    app.init();

})();


