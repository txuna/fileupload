const host = "127.0.0.1"
const port = 1234

const firstpage = 0

function init() {
    get_logs(firstpage).then(load_logs)
}

//처음 page는? 1
function get_logs(page) {
    console.log(page)
    return fetch(`http://${host}:${port}/statsLog/${page}`)
        .then(function(res) {
            return res.json()
        }).catch(function(err) {
            console.log(err)
        }).then(function(json) {
            return json
        })
}

function load_logs(logs) {
    //const table = document.querySelector(".logging__table")
    console.log(logs)
}

init()