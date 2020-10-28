
const mapDom = document.getElementById('map');
const spinnerDom = document.getElementById("spinner");

let map;
let latitude;
let longitude;
let myLocMarker;

document.addEventListener("DOMContentLoaded", function() {
    initMap();
});

function initMap() {
    if (localStorage.getItem("latitude") !== null && 
        localStorage.getItem("longitude") !== null) {
            latitude = localStorage.getItem("latitude");
            longitude = localStorage.getItem("longitude");
    } else {
        latitude = 37.566618;
        longitude = 126.978157;
    }

    var options = {
        center: new kakao.maps.LatLng(latitude, longitude),
        level: 5
    };
    map = new kakao.maps.Map(mapDom, options);
    gpsCheck(false);
}

function gpsCheck(flag) {
    navigator.geolocation.getCurrentPosition(function (pos) {
        localStorage.setItem("latitude", pos.coords.latitude);
        localStorage.setItem("longitude", pos.coords.longitude);

        if (flag) {
            myLocMarker.setMap(null);
            myLocMarker.position = new kakao.maps.LatLng(pos.coords.latitude, pos.coords.longitude);
            myLocMarker.setMap(map);
        } else {
            myLocMarker = new kakao.maps.CustomOverlay({
                position: new kakao.maps.LatLng(pos.coords.latitude, pos.coords.longitude),
                content: '<div class="my-location"><div class="my-location-h"><i class="fas fa-user"></i></div></div>',
                xAnchor: 0.5,
                yAnchor: 1.1,
                zIndex: 6
            });
        }
        myLocMarker.setMap(map);

        var moveLatLon = new kakao.maps.LatLng(pos.coords.latitude, pos.coords.longitude);
        map.panTo(moveLatLon);
    });
}

function getData(url, callBack) {
    var xhttp = new XMLHttpRequest();
    xhttp.onreadystatechange = callBack 
    xhttp.open("GET", url, true);
    xhttp.send();
}

function callBackNintendo() {
    if (this.readyState === 4) {
        if (this.status === 200) {
            var resp = JSON.parse(this.responseText);
            console.log(resp);
            spinnerDom.style.display = 'none';
        }
    }
};