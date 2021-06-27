var details = document.getElementById("details");
var currentView = "settings"

function selectView(view) {
    currentView = view;
    let i = document.getElementById("selected-controller").value;
    getSection("/list/"+currentView+"/", "details");
}

function selectController() {
    let i = document.getElementById("selected-controller").value;
    console.log("selected controller",i);
    fetch("/controller/"+i+"/")
    .then(response => {
        getSection("/list/"+currentView+"/", "details");
    })
    .catch(error => {
        console.error('Error:', error);
    });
}

function getSection(section,id) {
    fetch(section)
    .then(response => response.text())
    .then(text => {
        document.getElementById(id).innerHTML = text;
    })
    .catch(error => {
        console.error('Error:', error);
    });
}

function submitForm(frm) {
    let init = {
        method: 'POST', 
        body: new URLSearchParams(new FormData(frm)).toString(),
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded',
        },
        credentials: "include"
    };

    fetch(frm.action, init)
    .then(response => response.text())
    .then(text => {
        if (text.includes("error")) {
            alert(text);
            return;
        }
        let obj = JSON.parse(text);
        obj.forEach((e) => {
            document.getElementById(e.id).innerHTML = e.value;
        })
    })
    .catch(error => {
        console.error('Error:', error);
    });
}

function openEdit(key) {
    fetch("/edit/"+currentView+"/"+key+"/")
    .then(response => response.text())
    .then(text => {
        let dlg = document.getElementById("dialog")
        dlg.innerHTML = text;
        let frm = document.getElementById('settings-form')

        frm.addEventListener('submit', function(event) {
            event.preventDefault();
            dlg.style.display='none';
            submitForm(frm);
        }, false);

        dlg.style.display='block';
    })
    .catch(error => {
        console.error('Error:', error);
    });
}

function closeDialog() {
    document.getElementById("dialog").style.display='none';
}

function openOptions() {
    document.getElementById("options").style.display='block';
}

function closeOptions() {
    document.getElementById("options").style.display='none';
}

function applyOptions() {
    let theme = document.getElementById("selected-theme").value;
    let bug = document.getElementById("selected-icon").value;
    console.log(theme, bug);
    let init = {
        credentials: "include"
    };
    fetch("/options/"+theme+"/"+bug+"/", init);
    closeOptions();
}

function pop(doc) {
    let x = document.getElementById(doc);
    if (x.className.indexOf("w3-show") == -1) {
      x.className += " w3-show";
    } else { 
      x.className = x.className.replace(" w3-show", "");
    }
}

function openMenu(id) {
    let x = document.getElementById(id);
    x.className += " w3-show";
}
  
function closeMenu(id) {
    let x = document.getElementById(id);
    x.className = x.className.replace(" w3-show", "");
}
