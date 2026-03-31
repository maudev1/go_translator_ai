const helpers = {

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

