document.getElementById('convert-button').addEventListener('click', async () => {
    // get image
    let image = await getImage()
    if (image === null) {
        alert("no image")
        return
    } else {
        alert(111)
    }

    // convert it to boolean matrix
    // detect corners
    // strip fields

    // calculate pixel ratio coefficient

    // export to .xlsx
})

async function getImage() {
    let file = document.getElementById('image-upload').files[0]
    let resolve;

    // validate image
    if (file) {
        let image = new Image()
        // let url = window.URL || window.webkitURL
        image.src = URL.createObjectURL(file)

        image.addEventListener('load', () => {
            console.log(image)
            if (this.width) {
                console.log("seems to be an image, it has a width")
                
                resolve(image)
            } else {
                resolve(null)
            }
        })        
    }

    return new Promise((r) => {
        resolve = r
    })
}

function convertImageToMatrix(image) {
    let ctx = document.getElementById('canvas').getContext('2d')
    ctx.drawImage(image, 0, 0)
}