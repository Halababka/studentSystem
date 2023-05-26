import {ref} from "vue";
import {getCookie, logout, url} from "@/global";

export let userData = ref({
    first_name: "",
    last_name: "",
    middle_name: "",
    modal: false
});

export function getUserData() {
    const localeUserData = JSON.parse(localStorage.getItem("userData"));
    if (localeUserData !== null) {
        userData.value = localeUserData;
    } else {
        fetchUserData();
    }
}

function fetchUserData() {
    if (getCookie("ID") === undefined) logout();
    fetch(`${url}/user/${getCookie("ID")}`, {
        headers: {
            "Authorization": `Bearer ${getCookie("TOKEN")}`
        }
    }).then(response => {
        if (response.ok) return response.json();
        switch (response.status) {
            case 401:
                logout();
                break;
            case 404:
                // Сделать появление модального окна для ввода данных
                userData.value.modal = true
                break;
        }
    }).then(data => {
        if (data !== undefined) {
            userData.value = data;
            localStorage.setItem("userData", JSON.stringify(userData.value));
            // console.log("Finish fetch", userData.value);
        }
    }).catch(err => {
        console.error("Cannot fetch", err);
    });
    return {userData};
}