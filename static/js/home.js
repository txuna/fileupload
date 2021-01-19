//contextMenu를 킬때 어떤 파일을 대상으로 켰는지 확인하기 위한 전역변수
let contextFilename = undefined
let currentDir = "root"

const host = "127.0.0.1"
const port = 1234

const displayType = {
    "image": "far fa-file-image",
    "folder": "fas fa-folder",
    "text": "far fa-file-alt"
}

/* 
context-menu에 나타난 정보를 서버로 전송하여 값을 받는다.
값을 받을 때는 해당 메소드와 메소드에 해당하는 정보를 return 함
*/
function get_detail() {
    console.log(contextFilename)
    return fetch(`http://${host}:${port}/homeDetail/${contextFilename}`)
        .then(function(response) {
            return response.json()
        }).catch(function(err) {
            console.log(err)
        }).then(function(json) {
            return json
        })
}

//item을 클릭했을 때 어떤 행위를 할 것인가.
function click_Itemhandle(event) {
    if (event.target.id === "detail") {
        get_detail().then(function(result) {
            open_modal(result)
        })
    }
}

// contextmenu내부의 item에 대해 클릭 트리거 
function set_ItemclickEvent() {
    const items = document.querySelectorAll(".item")
    items.forEach(function(item) {
        item.addEventListener("click", click_Itemhandle)
    })
}

//dir들에 대해 contextmenu 이벤트 트리거를 한다.
function set_contextmenu() {
    const objects = document.querySelectorAll(".object")
    const contextElement = document.querySelector("#context-menu");
    objects.forEach(function(object) {
        object.addEventListener("contextmenu", function(event) {
            event.preventDefault(); //우클릭했을 때 기본적인 이벤트를 막는다. 
            contextElement.classList.remove("remove");
            contextElement.style.top = event.clientY + "px";
            contextElement.style.left = event.clientX + "px";
            contextElement.classList.add("active");
            //메뉴를 클릭한 아이템의 파일이름 가지고옴 - fullpath
            contextFilename = this.dataset.filepath
        });
    });
    window.addEventListener("click", function() {
        contextElement.classList.remove("active");
        contextElement.classList.add("remove");
    })
}

//모달창을 오픈하는 함수
function open_modal(filedetail) {
    const modal = document.querySelector("#detail-modal");

    modal.style.display = "block"

    const modal_content = document.querySelector(".modal__content")
    modal_content.innerHTML = `<span class="close">&times;</span>
    <div><span>파일 이름 : ${filedetail["filename"]}</span></div>
    <div><span>파일 사이즈 : ${filedetail["size"]}</span></div>
    <div><span>파일 경로 : ${filedetail["path"]}</span></div>
    <div><span>파일 유형 : ${filedetail["type"]}</span></div>
    <div><span>업로드 시간 : ${filedetail["uploadtime"]}</span></div>`

    const span = document.getElementsByClassName("close")[0];
    span.addEventListener("click", function() {
        modal.style.display = "none"
    })
}

//현재 경로를 기반으로 서버에 요청함.
function get_directory() {
    //? ?dir=/root
    return fetch(`http://${host}:${port}/homeDirectory/${currentDir}`)
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

//get_directory() 함수로 받아온 json array를 양식에 맞게 inline json -> 요청된 디렉토리를 프론트에 기입
function set_directory(dirInfo) {

    const current__path = document.querySelector("#current__path")
    const content_screen = document.querySelector(".content__screen")

    content_screen.innerHTML = ""
    current__path.innerText = `Current Path : ${dirInfo["path"]}`
    if (dirInfo["fs"] !== null) {
        dirInfo["fs"].forEach(function(e) {
            const div = document.createElement("div")
            const i = document.createElement("i")
            const span = document.createElement("span")

            div.classList.add("content", "object", e["display"])
            div.setAttribute("data-filepath", e["path"])
            i.classList.add(...displayType[e["display"]].split(" "))
            span.innerText = e["filename"]
            div.appendChild(i)
            div.appendChild(span)
            content_screen.appendChild(div)
        })
    }
    set_contextmenu()
    set_ItemclickEvent()
}

/*
folder나 file을 클릭했을 때 어떻게 할것인가 
folder를 누르면 folder에 적힌 dataset.filepath를 기반으로 homeDirectory fetch 
file을 누른다면 type파악하고 text라면 모달창으로 띄워주고 이미지라면 모달창으로 이미지 로드
*/
//this vs event.target => this는 이벤트가 처리되는 DOM 요소에 대한 참조이다. event.target은 이벤트를 시작한 요소를 나타낸다.
function click_dir(event) {
    event.preventDefault();
    const requestFilepath = this.dataset.filepath
    if (requestFilepath === undefined || requestFilepath === null) {
        alert("Incorrect Filepath...")
    } else {
        if (this.classList.contains("folder")) {
            console.log(requestFilepath)
            currentDir = requestFilepath
            get_directory().then(set_directory).then(set_clickdir)
        } else {
            alert("this is not folder")
        }
    }

}

function set_clickdir() {
    const objects = document.querySelectorAll(".object")
    console.log(objects)
    objects.forEach(function(e) {
        console.log(e)
        e.addEventListener("dblclick", click_dir)
    })
}

function init() {
    //get_directory()가 먼저 실행되어야 이벤트 등록가능 
    get_directory().then(set_directory).then(set_clickdir)
}

init()