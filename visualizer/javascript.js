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

const info = document.body.getBoundingClientRect()

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
        let maxX = 0;
        let minY = 100000;
        let maxY = 0;

        data.RoomS.forEach(room => {
            room.X < minX ? minX = room.X : room.X > maxX ? maxX = room.X : null
            room.Y < minY ? minY = room.Y : room.Y > maxY ? maxY = room.Y : null
            // console.log(room);
        });
        const ofset = 50
        const echelX =  (info.width / (maxX))
        const echelY =  (info.height / (maxY))
        console.log(echelX,echelY);
        ctx.clearRect(0, 0, canvas.clientWidth, canvas.clientHeight)
        data.RoomS.forEach(room => {
            ctx.beginPath()
            ctx.arc((room.X*echelX)+ofset, (room.Y*echelY)+ofset, 50, 0, Math.PI*2)
            ctx.fillStyle = "white"
            ctx.fill()
            // ctx.font = "50px Arial"
            // ctx.fillText(room.Name, room.X*70, room.Y*70, 1000)
            // ctx.strokeRect(room.X*70, room.Y*70-43, room.Name.length*40, 50)
        });
        console.log(minX, maxX, minY, maxY);
    });

