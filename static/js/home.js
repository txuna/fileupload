const objects = document.querySelectorAll(".object")

function init() {
    objects.forEach(function(object) {
        object.addEventListener("contextmenu", function(event) {
            event.preventDefault();
            const contextElement = document.querySelector("#context-menu");
            contextElement.classList.add("remove");
            contextElement.style.top = event.clientY + "px";
            contextElement.style.left = event.clientX + "px";
            contextElement.classList.remove("remove");
            contextElement.classList.add("active");

        });
    });
    window.addEventListener("click", function() {
        const contextElement = document.querySelector("#context-menu");
        contextElement.classList.remove("active");
        contextElement.classList.add("remove");
    })
}



init()