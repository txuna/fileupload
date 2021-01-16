const objects = document.querySelectorAll(".object")

function init() {
    objects.forEach(function(object) {
        object.addEventListener("contextmenu", function(event) {
            event.preventDefault();
            const contextElement = document.querySelector("#context-menu");
            contextElement.style.top = event.clientY + "px";
            contextElement.style.left = event.clientX + "px";
            contextElement.classList.add("active");

        });
    });
    window.addEventListener("click", function() {
        const contextElement = document.querySelector("#context-menu");
        contextElement.classList.remove("active");
    })
}

init()