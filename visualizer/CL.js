/*
    Cool DOwn Class
    Alann AS
*/

/**
 * 
 */
 class CooldDown {
    /**
    * @duration is time before func exec in miliseconds
    * @func is the function to start
    * @funcParam the params which need the function
    */
    constructor (duration, func, ...funcParam) {
        this.duration = duration
        this.func = func
        this.funcParam = funcParam
        this.isRuning = false
    }
    /**
    * Start cooldown if is not already start
    */
    start () {
        if (this.isRuning === true) return
        this.isRuning = true
        this.timeOut = setTimeout(()=>{this.func(...this.funcParam); this.isRuning = false}, this.duration)
    }
    /**
    * Interupt cooldown if is start whitout play the cooldown function
    */
    stop () {
        if (this.isRuning === false) return
        this.isRuning = false
        clearTimeout(this.timeOut)
    }
    /**
    * Interupt cooldown if is start and play the cooldown function
    */
    stopAndPlayFunc () {
        if (this.isRuning === false) return
        this.isRuning = false
        clearTimeout(this.timeOut)
        this.func(...this.funcParam)
    }
    /**
    * Stop and start the cooldown from begin whitout play the cooldown function
    */
    restart () {
        this.stop()
        this.start()
    }
}