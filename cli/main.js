var ws = new WebSocket("ws://localhost:8000/ws");

var map = {
    53: 1,  54: 2,  55: 3,   56: 4,
    84: 5,  89: 6,  85: 7,   73: 8,
    71: 9,  72: 10, 74: 11,  75: 12,
    66: 13, 78: 14, 188: 15, 186: 16,
}

ws.onmessage = function (event) {
    var msg = JSON.parse(event.data);
    if (msg.Type == 1) {
        document.getElementById("game").setAttribute('src', "data:image/png;base64," + msg.Screen);
    }
}

document.onkeydown = function (e) {
    e = e || window.event;
    if (map[e.keyCode]) {
        ws.send(JSON.stringify({
            Type: 3,
            Num: map[e.keyCode] - 1,
            State: true,
        }));
    }
}

document.getElementById("reset").onclick = function () {
    ws.send(JSON.stringify({
        Type: 4,
        Command: "Reset",
    }));
    ws.close(1000);
    location.reload(true);
}

document.getElementById("maze").onclick = function () {
    ws.send(JSON.stringify({
        Type: 5,
        Rom : "oh7CATIBohrQFHAEMEASAGAAcQQxIBIAEhiAQCAQIECAEA==",
    }));
    ws.close(1000);
    location.reload(true);
}

document.getElementById("tetris").onclick = function () {
    ws.send(JSON.stringify({
        Type: 5,
        Rom: "orQj5iK2cAHQETAlEgZx/9ARYBrQEWAlMQASDsRwRHASHMMDYB5hAyJc9RXQFD8BEjzQFHH/0BQjQBIc56EicuihIoTpoSKW4p4SUGYA9hX2BzYAEjzQFHEBEiqixPQeZgBDAWYEQwJmCEMDZgz2HgDu0BRw/yM0PwEA7tAUcAEjNADu0BRwASM0PwEA7tAUcP8jNADu0BRzAUMEYwAiXCM0PwEA7tAUc/9D/2MDIlwjNADugABnBWgGaQRhH2UQYgcA7kDgAABAwEAAAOBAAEBgQABAQGAAIOAAAMBAQAAA4IAAQEDAAADgIABgQEAAgOAAAEDAgADAYAAAQMCAAMBgAACAwEAAAGDAAIDAQAAAYMAAwMAAAMDAAADAwAAAwMAAAEBAQEAA8AAAQEBAQADwAADQFGY1dv82ABM4AO6itIwQPB58ATwefAE8HnwBI15LCiNykcAA7nEBE1BgG2sA0BE/AHsB0BFwATAlE2IA7mAb0BFwATAlE3SOEI3gfv9gG2sA0OE/ABOQ0OETlNDRewFwATAlE4ZLABOmff9+/z0BE4IjwD8BI8B6ASPAgKBtB4DSQAR1/kUCZQQA7qcA8lWoBPoz8mXwKW0ybgDd5X0F8Snd5X0F8ind5acA8mWitADuagBgGQDuNyM=",
    }));
    ws.close(1000);
    location.reload(true);
}

document.getElementById("invaders").onclick = function () {
    ws.send(JSON.stringify({
        Type: 5,
        Rom : "EiVTUEFDRSBJTlZBREVSUyB2MC45IEJ5IERhdmlkIFdJTlRFUmAAYQBiCKPT0BhxCPIeMSASLXAIYQAwQBItaQVsFW4AI4dgCvAV8AcwABJLI4d+ARJFZgBoHGkAagRrCmwEbTxuDwDgI2sjR/0VYATgnhJ9I2s4AHj/I2tgBuCeEosjazg5eAEjazYAEp9gBeCeEulmAWUbhICjz9RRo8/UUXX/Nf8SrWYAEunUUT8BEunUUWYAg0BzA4O1YviDImIIMwASySNzggZDCBLTMxAS1SNzggYzGBLdI3OCBkMgEuczKBLpI3M+ABMHeQZJGGkAagRrCmwEffRuDwDgI0cja/0VEm/3BzcAEm/9FSNHi6Q7EhMbfAJq/DsCEyN8AmoEI0c8GBJvAOCk02AUYQhiD9AfcAjyHjAsEzPwCgDgpvT+ZRIlo7f5HmEII1+BBiNfgQYjX4EGI1970ADugOCAEjAA28Z7DADuo89gHNgEAO4jR44jI0dgBfAY8BXwBzAAE38A7moAjeBrBOmhElemAv0e8GUw/xOlagBrBG0BbgETjaUA8B7bxnsIfQF6AToHE40A7jx+//+ZmX7//yQk537/PDx+24FCPH7/2xA4fP4AAH8APwB/AAAAAQEBAwMDAwAAPyAgICAgICAgPwgI/wAA/gD8AP4AAAB+QkJiYmJiAAD/AAAAAAAAAAD/AAD/AH0AQX0FfX0AAMLCxkRsKDgAAP8AAAAAAAAAAP8AAP8A9xAU9/cEBAAAfET+wsLCwgAA/wAAAAAAAAAA/wAA/wDvICjo6C8vAAD5hcXFxcX5AAD/AAAAAAAAAAD/AAD/AL4AIDAgvr4AAPcE54WFhPQAAP8AAAAAAAAAAP8AAP8AAH8APwB/AAAA7yjvAOBgbwAA/wAAAAAAAAAA/wAA/wAA/gD8AP4AAADAAMDAwMDAAAD8BAQEBAQEBAT8EBD/+YG5i5qa+gD6ipqam5n45iUl9DQ0NAAXFDQ3NibH31BQXNjY3wDfER8SGxnZfET+hoaG/IT+goL+/oDAwMD+/ILCwsL8/oD4wMD+/oDwwMDA/oC+hob+hob+hoaGEBAQEBAQGBgYSEh4nJCwwLCcgIDAwMD+7pKShoaG/oKGhoaGfIKGhoZ8/oL+wMDAfILCysR6/ob+kJyE/sD+AgL+/hAwMDAwgoLCwsL+goKC7jgQhoaWkpLugkQ4OESCgoL+MDAw/gIe8ID+AAAAAAYGAAAAYGDAAAAAAAAAGBgYGAAYfMYMGAAYAAD+/gAA/oKGhob+CAgIGBgY/gL+wMD+/gIeBgb+hMTE/gQE/oD+Bgb+wMDA/oL+/gICBgYGfET+hob+/oL+BgYGRP5ERP5EqKioqKioqGxaAAwYqDBOfgASGGZsqFpmVCRmAEhIGBKoBpCoEgB+MBKohDBOchhmqKioqKiokFR4qEh4bHKoEhhscmZUkKhyKhioME5+ABIYZmyoclSoWmYYfhhOcqhyKhgwZqgwTn4AbDBUTpyoqKioqKioSFR+GKiQVHhmqGwqMFqohDByKqjYqABOEqjkoqgAThKobCpUVHKohDByKqjenKhyKhioDFRIWnhyGGaochhCQmyocioAcqhyKhioME5+ABIYZmyoME4MZhgAbBiocioYMGaoHlRmDBicqCRUVBKoQngMPKiuqKioqKioqP8AAAAAAAAAAAAAAAAAAAA=",
    }));
    ws.close(1000);
    location.reload(true);
}

document.getElementById("pong").onclick = function () {
    ws.send(JSON.stringify({
        Type: 5,
        Rom : "IvZrDGw/bQyi6tq23NZuACLUZgNoAmBg8BXwBzAAEhrHF3cIaf+i8NZxourattzWYAHgoXv+YATgoXsCYB+LAtq2YAzgoX3+YA3goX0CYB+NAtzWovDWcYaEh5RgP4YCYR+HEkYAEnhGPxKCRx9p/0cAaQHWcRIqaAJjAYBwgLUSimj+YwqAcIDVPwESomECgBU/ARK6gBU/ARLIgBU/ARLCYCDwGCLUjjQi1GY+MwFmA2j+MwFoAhIWef9J/mn/Esh5AUkCaQFgBPAYdgFGQHb+Emyi8v4z8mXxKWQUZQDUVXQV8inUVQDugICAgICAgAAAAAAAayBsAKLq28F8ATwgEvxqAADu",
    }));
    ws.close(1000);
    location.reload(true);
}

document.getElementById("puzzle").onclick = function () {
    ws.send(JSON.stringify({
        Type: 5,
        Rom : "ahJrAWEQYgBgAKKw0SfwKTAA2rVxCHoIMTASJGEQcghqEnsIowDwHvBVcAEwEBIKahJrAWwAYv/ABnACIlJy/zIAEjhuAG4A8AoiUn4BfgESSISghbCGwDACEmRFARJkdfh2/DAIEnBFGRJwdQh2BDAGEnxEEhJ8dPh2/zAEEohEKhKIdAh2AaMA9h7wZYEAYACjAPYe8FWjAPwegBDwVfEp1FXatYpAi1CMYADu7l7+/v7+/v7+/g==",
    }));
    ws.close(1000);
    location.reload(true);
}

document.getElementById("close").onclick = function () {
    ws.close(1000);
}

document.getElementById("reload").onclick = function () {
    location.reload(true);
}

document.getElementById("nor").onclick = function () {
    ws.send(JSON.stringify({
        Type: 6,
        Clock: 500,
    }));
}

document.getElementById("tur").onclick = function () {
    ws.send(JSON.stringify({
        Type: 6,
        Clock: 750,
    }));
}

document.getElementById("utur").onclick = function () {
    ws.send(JSON.stringify({
        Type: 6,
        Clock: 1000,
    }));
}


document.getElementById("slowmo").onclick = function () {
    ws.send(JSON.stringify({
        Type: 6,
        Clock: 1,
    }));
}