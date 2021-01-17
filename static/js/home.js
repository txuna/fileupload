const objects = document.querySelectorAll(".object")
const contextElement = document.querySelector("#context-menu");
const items = document.querySelectorAll(".item")


//div가 아닌 i라면 상위 노드로 올라감

let current_object = {
    "filepath": "",
    "method": undefined
}

// context-menu에 나타난 정보를 서버로 전송하여 값을 받는다.
//값을 받을 때는 해당 메소드와 메소드에 해당하는 정보를 return 함
function fetch_Contextmenu(method) {
    current_object["method"] = method
    console.log(current_object)
    return {
        "method": "detail",
        "info": {
            "filename": "profile.jpg",
            "path": "/root/profile.jpg",
            "type": "jpeg",
            "size": "123213Byte",
            "UploadTime": "2021-01-03 13:03:12"
        }
    }
}

function click_Itemhandle(event) {
    const res = fetch_Contextmenu(event.target.id)
    if (res["method"] === "detail") {
        open_modal(res["info"])
    }
}

function set_ItemclickEvent() {
    items.forEach(function(item) {
        item.addEventListener("click", click_Itemhandle)
    })
}

function init() {
    objects.forEach(function(object) {
        object.addEventListener("contextmenu", function(event) {
            event.preventDefault(); //우클릭했을 때 기본적인 이벤트를 막는다.
            contextElement.classList.remove("remove");
            contextElement.style.top = event.clientY + "px";
            contextElement.style.left = event.clientX + "px";
            contextElement.classList.add("active");
            //child라면 
            current_object["filepath"] = this.dataset.filepath
        });
    });
    window.addEventListener("click", function() {
        contextElement.classList.remove("active");
        contextElement.classList.add("remove");
    })
}

function open_modal(filedetail) {
    const modal = document.querySelector("#detail-modal");

    modal.style.display = "block"

    const modal_content = document.querySelector(".modal__content")
    modal_content.innerHTML = `<span class="close">&times;</span>
    <div><span>파일 이름 : ${filedetail["filename"]}</span></div>
    <div><span>파일 사이즈 : ${filedetail["size"]}</span></div>
    <div><span>파일 경로 : ${filedetail["path"]}</span></div>
    <div><span>파일 유형 : ${filedetail["type"]}</span></div>
    <div><span>업로드 시간 : ${filedetail["UploadTime"]}</span></div>`

    const span = document.getElementsByClassName("close")[0];
    span.addEventListener("click", function() {
        modal.style.display = "none"
    })
}

init()
set_ItemclickEvent()