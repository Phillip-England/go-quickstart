

function qs(id) {
    return document.querySelector(id);
}

function qsa(id) {
    return document.querySelectorAll(id);
}

class Dom {

    static climbUntil(element, callback) {
        if (!element) {
            return qs('html')
        }
        if (callback(element) == true) {
            return element
        }
        return Dom.climbUntil(element.parentElement, callback)
    }

    static parseForm(form) {
        const formData = new FormData(form);
        const formObject = {};
        formData.forEach((value, key) => {
            formObject[key] = value;
        });
        let potentialErrElement = form.querySelector('.form-err');
        if (potentialErrElement) {
            formObject['err'] = potentialErrElement;
        }
        return formObject;
    }

}



class Animus {

    static defaultDuration = 100;

    static fadeIn(element, duration=Animus.defaultDuration, targetOpacity=1, display='flex') {
        element.style.display = display;
        element.style.opacity = 0;
        let start = null;
        function step(timestamp) {
            if (!start) start = timestamp;
            let progress = timestamp - start;
            let opacity = progress / duration * targetOpacity;
            element.style.opacity = Math.min(opacity, targetOpacity);
            if (progress < duration) {
                window.requestAnimationFrame(step);
            }
        }
        window.requestAnimationFrame(step);
    }

    static fadeOut(element, duration=Animus.defaultDuration, targetOpacity=0) {
        element.style.opacity = 1;
        let start = null;
        function step(timestamp) {
            if (!start) start = timestamp;
            let progress = timestamp - start;
            let opacity = 1 - progress / duration * (1 - targetOpacity);
            element.style.opacity = Math.max(opacity, targetOpacity);
            if (progress < duration) {
                window.requestAnimationFrame(step);
            } 
            // else {
            //     if (targetOpacity === 0) {
            //         element.style.display = 'none';
            //     }
            // }
        }
        window.requestAnimationFrame(step);
    }

}
