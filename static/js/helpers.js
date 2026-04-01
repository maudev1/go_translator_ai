const helpers = {

    init(){
        console.debug(`Helpers File Has Started`)
    },

    custom_alert(message, type) {
        Swal.fire({
            toast: true,
            position: "top-end",
            icon: type,
            title: message,
            showConfirmButton: false,
            timer: 1500
        });
    }


};


(function () {
    'use strict';
    helpers.init();
})();

