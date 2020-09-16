document.getElementById("spinner").style.display = 'block';
if (localStorage.getItem("switch_animal") === null) {
    localStorage.setItem("switch_animal", true);
    localStorage.setItem("switch_new", false);
}
var switch_new = localStorage.getItem("switch_new");
var switch_animal = localStorage.getItem("switch_animal");

var domSwitchNew = document.getElementById("switch_new");
var domSwitchAnimal = document.getElementById("switch_animal");

if (switch_new === 'true') {
    domSwitchNew.checked = true;
    var targetDom = document.getElementById("switch_new");
    var target = targetDom.parentNode.children[1];
    target.className = "fa fa-check-square-o"

    // Please vote -----------
    // target.className = "fa fa-square-o"
    // var elem = document.createElement("img");
    // elem.setAttribute("height", "15.2");
    // elem.setAttribute("width", "15.1");
    // elem.src = '/static/images/please_vote.svg';
    // target.appendChild(elem);
    // ---------
} else {
    domSwitchNew.checked = false;
    var targetDom = document.getElementById("switch_new");
    var target = targetDom.parentNode.children[1];
    target.className = "fa fa-square-o"

    // Please vote -----------
    // target.innerHTML = '';
    // target.className = "fa fa-square-o"
    // ---------
}
if (switch_animal === 'true') {
    domSwitchAnimal.checked = true;
    var targetDom = document.getElementById("switch_animal");
    var target = targetDom.parentNode.children[1];
    target.className = "fa fa-check-square-o"

    // Please vote -----------
    // target.className = "fa fa-square-o"
    // var elem = document.createElement("img");
    // elem.setAttribute("height", "15.2");
    // elem.setAttribute("width", "15.1");
    // elem.src = '/static/images/please_vote.svg';
    // target.appendChild(elem);
    // ---------
} else {
    domSwitchAnimal.checked = false;
    var targetDom = document.getElementById("switch_animal");
    var target = targetDom.parentNode.children[1];
    target.className = "fa fa-square-o"

    // Please vote -----------
    // target.innerHTML = '';
    // target.className = "fa fa-square-o"
    // ---------
}

document.getElementById("search").addEventListener('click', function (e) {
    if (e.target.nodeName == "BUTTON") {
        document.getElementById("search").classList.toggle('disable');
        var moveLatLon = new kakao.maps.LatLng(e.target.dataset.lat, e.target.dataset.lng);
        map.setLevel(3);
        map.panTo(moveLatLon);
    } else if (e.target.nodeName == "I") {
        document.getElementById("search").classList.toggle('disable');
    }
});

document.getElementById("filter").addEventListener('click', function (e) {
    if (e.target.type == "checkbox") {
        if (e.target.checked) {
            localStorage.setItem(e.target.id, true);
            location.reload();
        } else {
            localStorage.setItem(e.target.id, false);
            location.reload();
        }
    } else if (e.target.nodeName == "I") {
        document.getElementById("filter").classList.toggle('disable');
    }
});

var latitude
var longitude

if (localStorage.getItem("latitude") !== null && localStorage.getItem("longitude") !== null) {
    latitude = localStorage.getItem("latitude");
    longitude = localStorage.getItem("longitude");
} else {
    latitude = 37.566618;
    longitude = 126.978157;
}


var container = document.getElementById('map');
var options = {
    center: new kakao.maps.LatLng(latitude, longitude),
    level: 5
};

var map = new kakao.maps.Map(container, options);

var myLocMarker;

function gpsCheck(flag) {
    navigator.geolocation.getCurrentPosition(function (pos) {
        localStorage.setItem("latitude", pos.coords.latitude);
        localStorage.setItem("longitude", pos.coords.longitude);

        if (flag) {
            myLocMarker.setMap(null);
            myLocMarker.position = new kakao.maps.LatLng(pos.coords.latitude, pos.coords.longitude);
            myLocMarker.setMap(map);
        } else {
            myLocation = '<div class="my-location"><div class="my-location-h">Y</div></div>';
            myLocMarker = new kakao.maps.CustomOverlay({
                position: new kakao.maps.LatLng(pos.coords.latitude, pos.coords.longitude),
                content: myLocation,
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

function makeMark() {
    var searchHTML = "";
    for (var i = 0; i < DocsData.length; i++) {
        row = DocsData[i];
        // row.c[0]: {v: "게임샵"},이마트
        // row.c[1]: {v: "서울"},서울
        // row.c[2]: {v: "국전 한우리"},가양
        // row.c[3]: {v: "37.4847094,127.0177836"}
        // row.c[4]: {v: 0, f: "재고없음"}
        // row.c[5]: {v: "비고"}
        // row.c[6]: {v: "가능인원 동디션"}
        // row.c[7]: {v: "대기인원 동디션"}
        // row.c[8]: {v: "예상재고 동디션"}
        latlng = row.c[3].v.split(",")
        var position = new kakao.maps.LatLng(latlng[0], latlng[1]);

        var tmpClass = "stock-except";
        var optionClass = "";
        var tmpText = "확인중";
        var stockFlag = false;
        var waitableFlag = false;
        var indexN = 2;

        if (switch_new === 'true' && switch_animal === 'true') {
            if (row.c[4] != null && row.c[4].v != null) {
                if (row.c[4].v > 0) {
                    stockFlag = true;
                } else if (row.c[4].v == 0) {
                    tmpClass = "stock-no";
                    tmpText = "재고 없음";
                    indexN = 3;
                } else {
                    tmpClass = "stock-no";
                    tmpText = row.c[4].v;
                    indexN = 2;
                }
            }
            if (row.c[6] != null && row.c[6].v != null) {
                targetValue = row.c[6].v;
                var findFlag = targetValue.indexOf("명 이하");

                if (findFlag != -1) {
                    stockFlag = true;
                    waitableFlag = true;
                } else if (row.c[6].v == "인원마감") {
                    tmpClass = "stock-no end";
                    tmpText = "인원마감";
                    indexN = 4;
                } else if (row.c[6].v == "") {
                    tmpClass = "stock-except";
                    tmpText = "확인중";
                    indexN = 3;
                } else {
                    if (tmpText != "재고 없음") {
                        tmpClass = "stock-except";
                        tmpText = row.c[6].v;
                        indexN = 2;
                    }
                }
            }
        } else if (switch_new === 'true') {
            if (row.c[4] != null && row.c[4].v != null) {
                if (row.c[4].v > 0) {
                    stockFlag = true;
                } else if (row.c[4].v == 0) {
                    tmpClass = "stock-no";
                    tmpText = "재고 없음";
                    indexN = 3;
                } else {
                    tmpClass = "stock-no";
                    tmpText = row.c[4].v;
                    indexN = 2;
                }
            }
        } else if (switch_animal === 'true') {
            if (row.c[6] != null && row.c[6].v != null) {
                targetValue = row.c[6].v;
                var findFlag = targetValue.indexOf("명 이하");

                if (findFlag != -1) {
                    stockFlag = true;
                    waitableFlag = true;
                } else if (row.c[6].v == "인원마감") {
                    tmpClass = "stock-no end";
                    tmpText = "인원마감";
                    indexN = 4;
                } else if (row.c[6].v == "") {
                    tmpClass = "stock-except";
                    tmpText = "확인중";
                    indexN = 3;
                } else {
                    tmpClass = "stock-except";
                    tmpText = row.c[6].v;
                    indexN = 2;
                }
            }
        } else {
            document.getElementById("spinner").style.display = 'none';
            return
        }

        if (stockFlag) {
            tmpClass = "stock-yes";
            if (waitableFlag) {
                indexN = 6;
                var able = "";
                able = row.c[6].v
                tmpText = "<span>" + able.split(" ")[0] + '/</span><span class="pre">' +
                    row.c[8].v + "</span>";
            } else {
                indexN = 5;
                tmpText = "재고 있음";
            }
        }

        var titleStr = "";
        if (row.c[0] != null && row.c[0].v != null) {
            titleStr += row.c[0].v;
            if (row.c[0].v == "이마트" || row.c[0].v == "트레이더스" || row.c[0].v == "일렉트로마트") {
                tmpClass += " emart"
            } else if (row.c[0].v == "롯데마트" || row.c[0].v == "토이저러스") {
                tmpClass += " lotte"
            }
        }
        if (row.c[2] != null && row.c[2].v != null) {
            if (titleStr != "") {
                titleStr += " ";
            }
            titleStr += row.c[2].v;
        }

        var content = '<div class="overlay_info ' + tmpClass + '" title="' + titleStr +
            '" data-num="' + i + '" onclick="getInfoDetail(' + i + ')"><span class="address">' + tmpText + '</span></div>';

        var mapCustomOverlay = new kakao.maps.CustomOverlay({
            position: position,
            content: content,
            xAnchor: 0.5,
            yAnchor: 1.1,
            zIndex: indexN
        });
        mapCustomOverlay.setMap(map);
        searchHTML += '<li><button data-lat="' + latlng[0] +
            '" data-lng="' + latlng[1] + '">' + titleStr + '</button></li>';
    }
    document.getElementById("search_list").innerHTML = searchHTML;
    document.getElementById("spinner").style.display = 'none';
}

getSwitchDocsData();
gpsCheck(false);
viewNotice(false);
// loadpageViewNotice();

var DocsData;

function getSwitchDocsData() {
    var xhttp = new XMLHttpRequest();
    xhttp.onreadystatechange = function () {
        if (this.readyState === 4) {
            if (this.status === 401) {
                // sessiong check
            } else if (this.status === 200) {
                d = this.responseText;
                d = d.replace("/*O_o*/", "");
                d = d.replace(/(^\ngoogle\.visualization\.Query\.setResponse\(|\);$)/g, '');
                var dataOBJ = JSON.parse(d);
                DocsData = dataOBJ.table.rows;
                data = dataOBJ.table.rows;
                makeMark();
                // makeArray();
            } else {
                return;
            }
        }
    };
    xhttp.open("GET", "https://docs.google.com/spreadsheets/d/1cJjX1AgTGRW6DF-ibDTKkd87wZ85wIOFi91Zpokmg3U/gviz/tq?tq=SELECT+C%2cB%2cD%2cF%2cJ%2cK%2cG%2cH%2cI+WHERE+F!%3d%22%22&sheet=동디션_매장제보", true);
    // xhttp.open("GET", "https://docs.google.com/spreadsheets/d/1YEnPNIe25zmRS3URQYVkYR791cVK7TYIpgA0oAygwhA/gviz/tq?tq=SELECT+C%2cB%2cD%2cF%2cJ%2cK%2cG%2cH%2cI+WHERE+F!%3d%22%22", true);
    xhttp.send();
}

function getInfoDetail(num) {
    row = DocsData[num];

    var titleStr = "";
    if (row.c[0] != null && row.c[0].v != null) {
        titleStr += row.c[0].v;
    }
    if (row.c[2] != null && row.c[2].v != null) {
        titleStr += " " + row.c[2].v;
    }

    var html = '<div class="info-detail">';
    if (switch_animal === 'true') {
        html += '<p style="margin-bottom: 8px;">동물의 숲 에디션 재고</p><table border="1"><tr>' +
            '<td>가능 인원</td><td>대기 인원</td><td>예상 재고</td></tr><tr>';
        if (row.c[6] != null && row.c[6].v != null && row.c[6].v != 0) {
            html += '<td>' + row.c[6].v + '</td>';
        } else {
            html += '<td>재고 없음</td>';
        }
        if (row.c[7] != null && row.c[7].v != null && row.c[7].v != 0) {
            html += '<td>' + row.c[7].v + '명 이상</td>';
        } else {
            html += '<td>제보 없음</td>';
        }
        if (row.c[8] != null && row.c[8].v != null && row.c[8].v != 0) {
            html += '<td>' + row.c[8].v + '개 이하</td>';
        } else {
            html += '<td>재고 없음</td>';
        }
        html += '</tr></table>';
    }
    if (switch_new === 'true') {
        html += '<p>닌텐도 신형 재고 : '
        if (row.c[4] != null && row.c[4].v != null && row.c[4].v != 0) {
            html += row.c[4].v + '개</p>';
        } else if (row.c[4] == null || row.c[4].v == null) {
            html += '확인 중</p>';
        } else {
            if (row.c[4].v == 0) {
                html += '재고 없음</p>';
            } else {
                html += '확인 중</p>';
            }
        }
    }
    if (row.c[5] != null && row.c[5].v != null) {
        html += '<p>비고 : ' + row.c[5].v + '</p>';
    }

    swal({
        title: titleStr,
        width: 400,
        html: html
    });
}

$('#search-market').keyup(function (e) {
    searchMarket('search-market');
});

function searchMarket(target) {
    var targetDom = document.getElementById(target);
    var keyWord = document.getElementById(target).value;
    var targetList = targetDom.parentNode.parentNode.children[1];

    if (keyWord == "") {
        cleanSearch(target);
    } else {
        list = targetList.children;
        if (list != null && list != undefined) {
            if (list.length != 0) {
                for (let i = 0; i < list.length; i++) {
                    targetValue = list[i].children[0].innerText;
                    var findFlag = targetValue.indexOf(keyWord);
                    if (findFlag != -1) {
                        list[i].style.display = "block";
                    } else {
                        list[i].style.display = "none";
                    }
                }
            }
        }
    }
}

function cleanSearch(target) {
    var targetDom = document.getElementById(target);
    var targetList = targetDom.parentNode.parentNode.children[1];

    document.getElementById(target).value = "";
    list = targetList.children;
    if (list != null && list != undefined) {
        if (list.length != 0) {
            for (let i = 0; i < list.length; i++) {
                list[i].style.display = "block";
            }
        }
    }
}

function viewNotice(flag) {
    var xhttp = new XMLHttpRequest();
    xhttp.onreadystatechange = function () {
        if (this.readyState === 4) {
            if (this.status === 401) {
                // sessiong check
            } else if (this.status === 200) {
                d = this.responseText;
                d = d.replace("/*O_o*/", "");
                d = d.replace(/(^\ngoogle\.visualization\.Query\.setResponse\(|\);$)/g, '');
                var dataOBJ = JSON.parse(d);
                noticeData = dataOBJ.table.rows;

                if (!flag) {
                    if (localStorage.getItem("view_notice") != null) {
                        sdStr = localStorage.getItem("view_notice");
                        noticeCnt = localStorage.getItem("view_notice_cnt");
                        nd = new Date();
                        ndStr = nd.getFullYear() + "/" + (nd.getMonth() + 1) + "/" + nd.getDate();
                        if (sdStr == ndStr && noticeCnt == noticeData.length) {
                            return
                        }
                    }
                }

                var html = "";
                for (var i = 0; i < noticeData.length; i++) {
                    row = noticeData[i];
                    // row.c[0]: {v: "Date(2020,3,6)", f: "2020. 4. 6"}
                    // row.c[1]: {"v": [21,27,26,972],"f": "오후 9:27:27"}
                    // row.c[2]: {v: "공지 내용"}

                    if (row.c[0] != null && row.c[0].v != null && row.c[0].f != null) {
                        if (row.c[1] != null && row.c[1].v != null) {
                            var notDate;
                            eval("notDate = new " + row.c[0].v + ";")
                            notDate.setHours(row.c[1].v[0]);
                            notDate.setMinutes(row.c[1].v[1]);

                            var currentDate = notDate.getFullYear() + "/" + (notDate.getMonth() + 1) + "/" + notDate.getDate() + " ";
                            currentDate += row.c[1].v[0] + ":" + row.c[1].v[1] + "";

                            html = '<div class="notice-item"><p>' + currentDate + '</p>'
                                + row.c[2].v + '</div>' + html;
                        }
                    }
                }

                swal({
                    title: "공지 사항",
                    width: 400,
                    html: '<div class="notice-main">' + html + '</div>'
                }).then(function () {
                    nd = new Date();
                    localStorage.setItem("view_notice", nd.getFullYear() + "/" + (nd.getMonth() + 1) + "/" + nd.getDate());
                    localStorage.setItem("view_notice_cnt", noticeData.length);
                });
            } else {
                return;
            }
        }
    };
    xhttp.open("GET", "https://docs.google.com/spreadsheets/d/1tuhm4dUc-YlhwhKNgKsHtg4IyvbtzJdnG0LaZmk4ZrI/gviz/tq?tq=SELECT+B%2cC%2cD+WHERE+D!%3d%22%22", true);
    xhttp.send();
}

// const res = Array();
// var data;
// // JSON.stringify(res)
// j = 0;
// console.log("3");
// function makeArray() {
//     setTimeout(function () {
//         row = data[j];
//         var tmpAry = new Array();
//         if (row.c[1] == null || row.c[1].v == null || row.c[2] == null || row.c[2].v == null) {
//             res.push(tmpAry)
//         } else {
//             t = row.c[1].v
//             tmpAry[0] = t.replace("-", "+");
//             tmpAry[1] = row.c[2].v
//             tmpAry[2] = row.c[0].v
//             tmpAry[3] = "";
//             res.push(tmpAry)

//             if (row.c[0].v != "게임샵") {
//                 urlparm = tmpAry.join('+');
//                 // console.log(urlparm);
//                 getlatlng(urlparm, j);
//             } else {
//                 urlparm = tmpAry[0] + tmpAry[1]
//                 // console.log(urlparm);
//                 getlatlng(urlparm, j);
//             }
//         }


//         j++;
//         if (j < data.length) {
//             console.log(j);
//             makeArray();
//         }
//     }, 20)
// }

// function getlatlng(urlparm, i) {
//     url = "https://maps.googleapis.com/maps/api/geocode/json?address=" + urlparm + "&key="
//     var latlan
//     var xhttp = new XMLHttpRequest();
//     xhttp.onreadystatechange = function () {
//         if (this.readyState === 4) {
//             if (this.status === 401) {
//                 // sessiong check
//             } else if (this.status === 200) {
//                 var dataOBJ = JSON.parse(this.responseText);
//                 // dataOBJ.results[0].geometry.location.lat;
//                 // dataOBJ.results[0].geometry.location.lng;

//                 if (dataOBJ.results[0] != undefined) {
//                     latlan = dataOBJ.results[0].geometry.location.lat + "," + dataOBJ.results[0].geometry.location.lng;
//                     res[i][3] = latlan;
//                 } else {
//                     console.log(urlparm);
//                     console.log(dataOBJ);
//                     // getlatlng(urlparm, i);
//                 }
//             } else {
//                 console.log(urlparm);
//             }
//         }
//     }
//     xhttp.open("GET", url, true);
//     xhttp.send();
// }