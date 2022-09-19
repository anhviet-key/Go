/** @format */

const $ = document.querySelector.bind(document);
const $$ = document.querySelectorAll.bind(document);
const $$$ = document.getElementById.bind(document);
const url = "http://localhost:8080/api/todo/";
const bool = true;
const app = {
  bool: true,
  active: function () {
    $$(".circle").forEach(function (e) {
      e.addEventListener("click", function () {
        e.classList.toggle("active");
      });
    });
    $$$("input").addEventListener("keypress", (even) => {
      if (even.key == "Enter") {
        if ($$$("input").value == "") {
        } else {
          $$$("submitForm").removeAttribute("onsubmit");
        }
        return;
      }
    });
  },
  delete: function () {
    $$(".linkDel").forEach(function (e) {
      e.addEventListener("click", function (v) {
        const id = e.getAttribute("data");
        var options = {
          method: "DELETE",
          headers: {
            "Content-Type": "application/json",
          },
        };
        fetch(url + id, options).then(function (response) {
          response.json();
          location.reload();
        });
      });
    });
  },
  update: function () {
    $$(".linkEdit").forEach(function (e) {
      e.addEventListener("click", function (v) {
        e.parentNode.parentNode
          .querySelector(".formEdit")
          .classList.toggle("block");
        const id = e.getAttribute("data");
        const context = e.parentNode.parentNode.querySelector(".formEdit");
        context.addEventListener("keypress", (even) => {
          if (even.key == "Enter") {
            if (context.value == "") {
              return;
            } else {
              console.log("old key pressed");
              var options = {
                method: "PUT",
                body: JSON.stringify({
                  Title: context.value,
                  Status: bool,
                }),
                headers: {
                  "Content-Type": "application/json",
                },
              };
              fetch(url + id, options)
                .then(function (response) {
                  response.json();
                  location.reload();
                })
                .catch(function (error) {
                  alert("error");
                });
              return;
            }
          }
        });
      });
    });
  },
  success: function () {
    $$(".circle").forEach(function (e) {
      e.addEventListener("click", function () {
        const id = e.getAttribute("data");
        e.classList.toggle("active") ? (bool = true) : (bool = false);
        const context = e.parentNode.querySelector(".txt");
        var options = {
          method: "PUT",
          body: JSON.stringify({
            Title: context.innerText,
            Status: bool,
          }),
          headers: {
            "Content-Type": "application/json",
          },
        };
        fetch(url + id, options)
          .then(function (response) {
            response.json();
            location.reload();
          })
          .catch(function (error) {
            alert("error");
          });
        return;
      });
    });
  },
  start: function () {
    this.active();
    this.delete();
    this.update();
    this.success();
  },
};

app.start();
