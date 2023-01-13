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

let wait = true
let step = -1
let isUpdate = false
const travelTime = 2000
const waitTime = 5000
const CoolDownTravel = new CooldDown(travelTime, ()=> wait = true);
const CoolDownWaitTravel = new CooldDown(waitTime, ()=> {wait = false; step++});

/**
 * 
 * @param {Array} tableObj 
 * @param {string} objName 
 */
function getRoomByName(tableObj, objName) {
    return tableObj.find(v => v.Name === objName)
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

class Ant {
    constructor(name, x, y) {
        this.name = name
        this.x = x
        this.y = y
        this.xSpeed = 0
        this.ySpeed = 0
    }
    /**
     * 
     * @param {Array} ants 
     */
    static getByName(ants, nameInt) {
        return ants.find(v => v.name == nameInt)
    }
}


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
        
        tableAnt = []
        for (let i = 0; i < data.NbAnt; i++) {
            tableAnt.push(new Ant(i+1, data.StartRoom.X*echelX+d, data.StartRoom.Y*echelY+d))
        }
        console.log(tableAnt);
        requestAnimationFrame(loop = ()=>{
            CoolDownWaitTravel.start()
            if (wait || data.Steps.length <= step ) {
                requestAnimationFrame(loop)
                isUpdate = false
                return
            }
            ctx.clearRect(0, 0, canvas.clientWidth, canvas.clientHeight)
            data.RoomS.forEach(room => {
                room.Link.forEach((v)=>{
                    let l = getRoomByName(data.RoomS, v)
                    // console.log(l);
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
    
    
            // console.log(echelX,echelY);
            data.RoomS.forEach(room => {
                ctx.beginPath()
                ctx.arc((room.X*echelX)+ofset, (room.Y*echelY)+ofset, d, 0, Math.PI*2)
                ctx.fillStyle = "white"
                ctx.fill()
                ctx.fillStyle = "red"
                // const fontSize = Math.floor(2 * d / (1.8 * ctx.measureText('M').width));
                // const widthText = ctx.measureText(room.Name).width
                // ctx.font = fontSize + "px Arial"
    
                var fontSize = 48
                ctx.font = "10px Arial";
                // Mesure la largeur du texte
                var textWidth = ctx.measureText(room.Name).width;
                // Calcule la taille de police adaptée
                if (textWidth > 14) {
                    var fontSize = Math.floor((d / textWidth * 18));
                }
                // Définit la taille de police
                ctx.font = fontSize + "px Arial";
                // console.log(room.Name, textWidth, fontSize, 2*d-ctx.measureText(room.Name).width);
                var w = 2*d-ctx.measureText(room.Name).width
                ctx.fillText(room.Name, (room.X*echelX)+w/2 , (room.Y*echelY)+(d*2)-(d*2-fontSize*0.7)/2, d*2)
                // ctx.strokeRect((room.X*echelX)+ofset , (room.Y*echelY)+ofset, room.Name.length*40, 50)
            });

            data.Steps[step].Ants.forEach((antName, i)=>{
                /**
                 * @type {Ant}
                 */
                const ant = Ant.getByName(tableAnt, antName.Name)
                // console.log(ant);
                if (!isUpdate) {
                    const r = data.Steps[step].Paths[i]
                    const speeds = GetSpeedsToGo(ant.x, ant.y, r.X*echelX+d, r.Y*echelY+d, travelTime)
                    ant.xSpeed = speeds[0]*16,7
                    ant.ySpeed = speeds[1]*16,7
                    
                    console.log("speeds", ant.xSpeed, ant.ySpeed);
                }
                CoolDownTravel.start()
                    ctx.beginPath()
                    ctx.arc(ant.x, ant.y, 20, 0, Math.PI*2)
                    ctx.fillStyle = "blue"
                    ctx.fill()
                ant.x += ant.xSpeed
                ant.y += ant.ySpeed
            })
            isUpdate = true
            requestAnimationFrame(loop)
        })
    });





function GetSpeedsToGo(xStart, yStart, xEnd, yEnd, travelTime) {
    const speedX = (xEnd - xStart) / travelTime;
    const speedY = (yEnd - yStart) / travelTime;
    return [speedX, speedY]
}