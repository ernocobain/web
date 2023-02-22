let hamburger = document.querySelector("#hamburger")
let nav = document.querySelector("nav")
let activeMenu = document.querySelectorAll(".hamburger")

hamburger.addEventListener("click", function() {
    nav.classList.toggle("hidden")
    for (let index = 0; index < activeMenu.length; index++) {
        activeMenu[index].classList.toggle("active")
    }
})