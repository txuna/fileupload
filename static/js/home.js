//div가 아닌 i라면 상위 노드로 올라감

let current_object = {
    "filepath": "",
    "method": undefined
}

//contextMenu를 킬때 어떤 파일을 대상으로 켰는지 확인하기 위한 전역변수
let contextFilename = undefined

const host = "127.0.0.1"
const port = 1234

// context-menu에 나타난 정보를 서버로 전송하여 값을 받는다.
//값을 받을 때는 해당 메소드와 메소드에 해당하는 정보를 return 함
function get_detail() {
    return fetch(`http://${host}:${port}/homeDetail?filepath=${contextFilename}`)
        .then(function(response) {
            return response.json()
        }).catch(function(err) {
            console.log(err)
        }).then(function(json) {
            return json
        })
}

function click_Itemhandle(event) {
    if (event.target.id === "detail") {
        get_detail().then(function(result) {
            open_modal(result)
        })
    }
}

function set_ItemclickEvent() {
    const items = document.querySelectorAll(".item")
    items.forEach(function(item) {
        item.addEventListener("click", click_Itemhandle)
    })
}

function set_contextmenu() {
    const objects = document.querySelectorAll(".object")
    console.log(objects)
    const contextElement = document.querySelector("#context-menu");
    objects.forEach(function(object) {
        object.addEventListener("contextmenu", function(event) {
            event.preventDefault(); //우클릭했을 때 기본적인 이벤트를 막는다. 
            console.log(object)
            contextElement.classList.remove("remove");
            contextElement.style.top = event.clientY + "px";
            contextElement.style.left = event.clientX + "px";
            contextElement.classList.add("active");
            //child라면 
            //current_object["filepath"] = this.dataset.filepath 
            contextFilename = this.dataset.filepath
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
    <div><span>파일 이름 : ${filedetail["name"]}</span></div>
    <div><span>파일 사이즈 : ${filedetail["size"]}</span></div>
    <div><span>파일 경로 : ${filedetail["path"]}</span></div>
    <div><span>파일 유형 : ${filedetail["type"]}</span></div>
    <div><span>업로드 시간 : ${filedetail["uploadtime"]}</span></div>`

    const span = document.getElementsByClassName("close")[0];
    span.addEventListener("click", function() {
        modal.style.display = "none"
    })
}

function get_directory() {
    //? ?dir=/root
    const dir = "/root"
    return fetch(`http://${host}:${port}/homeDirectory?dir=${dir}`)
        .then(function(response) {
            return response.json()
        })
        .catch(function(err) {
            console.log(err)
        })
        .then(function(json) {
            return json
                //console.log("result : " + JSON.stringify(json))
        })
}
const displayType = {
        "image": "far fa-file-image",
        "folder": "fas fa-folder",
        "text": "far fa-file-alt"
    }
    //get_directory() 함수로 받아온 json array를 양식에 맞게 inline json

function set_directory(filejson) {
    const content_screen = document.querySelector(".content__screen")
    console.log(filejson)
    filejson.forEach(function(e) {
        const div = document.createElement("div")
        const i = document.createElement("i")
        const span = document.createElement("span")

        div.classList.add("content", "object", e["display"])
        div.setAttribute("data-filepath", e["path"])
        i.classList.add(...displayType[e["display"]].split(" "))
        span.innerText = e["name"]
        div.appendChild(i)
        div.appendChild(span)
        content_screen.appendChild(div)
    })
    set_contextmenu()
    set_ItemclickEvent()
}

function init() {
    //get_directory()가 먼저 실행되어야 이벤트 등록가능 
    get_directory().then(set_directory)
}

init()