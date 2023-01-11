console.log("Js Start");

//GET CANVAS
const canvas = document.getElementById("canvas")

// canvas.style.height = document.body.clientHeight + "px"
// canvas.style.width = document.body.clientWidth + "px"
canvas.style.background = document.body.style.background = "black"
/**
 * @type {CanvasRenderingContext2D}
 */
const ctx = canvas.getContext('2d');
ctx.clearRect(0, 0, canvas.clientWidth, canvas.clientHeight)
const info = canvas.getBoundingClientRect()

/**
 * 
 * @param {Array} tableObj 
 * @param {string} objName 
 */
function getRoomByName(tableObj, objName) {
    
    return tableObj.filter(v => {
        return v.Name == objName
    })
}

/*
ctx.fillStyle = 'red';
ctx.fillRect(0, 0, 100, 100);

ctx.fillStyle = 'blue';
ctx.fillRect(100, 100, 100, 100);

ctx.fillStyle = 'green';
ctx.arc(150, 150, 50, 0, Math.PI * 2);
ctx.fill();
*/



// FETCH AND UNMARCHAL DATA____
fetch('../ant.json')
    .then(response => response.text())
    .then(text => {
        const data = JSON.parse(text);
        console.log(data);
        //DRAW each Room and link
        
        let minX = 100000;
        let maxX = 10;
        let minY = 100000;
        let maxY = 5;
        const d = 60
        const ofset = d


        data.RoomS.forEach(room => {
            console.log(room.X, room.Y);
            room.X < minX ? minX = room.X : room.X > maxX ? maxX = room.X : null
            room.Y < minY ? minY = room.Y : room.Y > maxY ? maxY = room.Y : null
            // console.log(room);
        });
        const echelX =  ((info.width-(2*ofset)) / (maxX))
        const echelY =  ((info.height-(2*ofset)) / (maxY))
        data.RoomS.forEach(room => {

            room.Link.forEach((v)=>{
                let l = getRoomByName(data.RoomS, v)[0]
                // if (l.isLinke == undefined || !l.isLinke.includes(l.Name)) {
                    // l.isLinke = Array()
                    // l.isLinke.push(room.Name)
                    ctx.beginPath();
                    ctx.moveTo(room.X*echelX+ofset, room.Y*echelY+ofset);
                    ctx.lineTo(l.X*echelX+ofset, l.Y*echelY+ofset);
                    ctx.strokeStyle = "white";
                    ctx.lineWidth = 10
                    ctx.stroke();
                // }
            })
        });


        console.log(echelX,echelY);
        data.RoomS.forEach(room => {
            ctx.beginPath()
            ctx.arc((room.X*echelX)+ofset, (room.Y*echelY)+ofset, d, 0, Math.PI*2)
            ctx.fillStyle = "white"
            ctx.fill()
            ctx.fillStyle = "red"
            // const fontSize = Math.floor(2 * d / (1.8 * ctx.measureText('M').width));
            // const widthText = ctx.measureText(room.Name).width
            // ctx.font = fontSize + "px Arial"

            var fontSize = 36
            ctx.font = "10px Arial";
            // Mesure la largeur du texte
            var textWidth = ctx.measureText(room.Name).width;
            // Calcule la taille de police adaptée
            if (textWidth > 14) {
                var fontSize = Math.floor((d / textWidth * 18));
            }
            // Définit la taille de police
            ctx.font = fontSize + "px Arial";
            console.log(room.Name, textWidth, fontSize, 2*d-ctx.measureText(room.Name).width);
            var w = 2*d-ctx.measureText(room.Name).width
            ctx.fillText(room.Name, (room.X*echelX)+w/2 , (room.Y*echelY)+(d*2)-(d*2-fontSize*0.7)/2, d*2)
            // ctx.strokeRect((room.X*echelX)+ofset , (room.Y*echelY)+ofset, room.Name.length*40, 50)

        });


        console.log(minX, maxX, minY, maxY);
    });

