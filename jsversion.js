// var input
// import { createRequire } from 'module';
// const require = createRequire(import.meta.url);

// 给当前上传文件button添加监听事件，一般初始化的时候就可以执行
function getInputTxt(){
  this.txtDom = document.getElementById('readTxt');
  this.txtDom.addEventListener('change', this.updateTxt, false);
};

function updateTxt(e){
  let selectedFile = document.getElementById('readTxt').files[0];
  let reader = new FileReader();
  reader.readAsText(selectedFile);
  reader.onload = oFREvent => {
      let pointsTxt = oFREvent.target.result;
      console.log('pointsTxt', pointsTxt)
  };
  e.target.value = null;
}
function removeTxtDomListener() {
  this.txtDom.removeEventListener('change', this.updateTxt, false);
}


const findTXTfile = fs.readdirSync('./txtFolder');
let txtFileList=[];
let excelFileList=[];
/////////////////////////////////////////////////////////////////
function checkFileAPI() {
    if (window.File && window.FileReader && window.FileList && window.Blob) {
        reader = new FileReader();
        return true; 
    } else {
        alert('The File APIs are not fully supported by your browser. Fallback required.');
        return false;
    }
}

/**
 * read text input
 */
function readText(filePath) {
    var output = ""; //placeholder for text output
    if(filePath.files && filePath.files[0]) {           
        reader.onload = function (e) {
            output = e.target.result;
            displayContents(output);
        };//end onload()
        reader.readAsText(filePath.files[0]);
    }//end if html5 filelist support
    else if(ActiveXObject && filePath) { //fallback to IE 6-8 support via ActiveX
        try {
            reader = new ActiveXObject("Scripting.FileSystemObject");
            var file = reader.OpenTextFile(filePath, 1); //ActiveX File Object
            output = file.ReadAll(); //text contents of file
            file.Close(); //close file "input stream"
            displayContents(output);
        } catch (e) {
            if (e.number == -2146827859) {
                alert('Unable to access local files due to browser security settings. ' + 
                 'To overcome this, go to Tools->Internet Options->Security->Custom Level. ' + 
                 'Find the setting for "Initialize and script ActiveX controls not marked as safe" and change it to "Enable" or "Prompt"'); 
            }
        }       
    }
    else { //this is where you could fallback to Java Applet, Flash or similar
        return false;
    }       
    return true;
}   

/**
 * display content using a basic HTML replacement
 */
function displayContents(txt) {
    var el = document.getElementById('main'); 
    el.innerHTML = txt; //display output in DOM
} 
//////////////////////////////////////////////////////////////////

async function getTxtfile(){
    for(x=0;x<findTXTfile.length;x++){
        if(findTXTfile[x].includes('.txt')){
           txtFileList.push(findTXTfile[x])
            document.querySelector('.selector').innerHTML+=`<option>${findTXTfile[x]}</option>`
        }
    }
}
document.querySelector('.export_button').innerHTML+=1321321
window.onload(getTxtfile())

console.log(txtFileList[0])

// console.log(file)
// console.log(text)

// console.log(textByLine)

// console.log(formatmessagefifteen)
formatmessage15array = textByLine[15]
  .substring(textByLine[15].indexOf("{") + 1, textByLine[15].lastIndexOf("}"))
  .split(";");
var product = formatmessage15array[0].split(":")[1];
var quantity = formatmessage15array[1].split(":")[1];
var price = formatmessage15array[2].split(":")[1];
var date = textByLine[15].substring(
  textByLine[15].indexOf("[") + 1,
  textByLine[15].lastIndexOf("]")
);
var name = textByLine[15].split(":")[2].split("]")[1];

function getProduct(message) {
  if (
    message
      .substring(message.indexOf("{") + 1, message.lastIndexOf("}"))
      .split(";")[0].length !== 0
  ) {
  }
}

var fs = require("fs");
var text = fs.readFileSync("./txtFolder/chat.txt", "utf-8");
var textByLine = text.split("\n");
var validRecord = [];
for (x = 0; x < textByLine.length; x++) {
  if (
    textByLine[x]
      .substring(textByLine[x].indexOf("{") + 1, textByLine[x].lastIndexOf("}"))
      .split(";")[0].length !== 0
  ) {
    validRecord.push({
      Date: textByLine[x].substring(
        textByLine[x].indexOf("[") + 1,
        textByLine[x].lastIndexOf("]")
      ),
      Name: textByLine[x].split(":")[2].split("]")[1],
      Product: textByLine[x]
        .substring(
          textByLine[x].indexOf("{") + 1,
          textByLine[x].lastIndexOf("}")
        )
        .split(";")[0]
        .split(":")[1],
      Quantity: textByLine[x]
        .substring(
          textByLine[x].indexOf("{") + 1,
          textByLine[x].lastIndexOf("}")
        )
        .split(";")[2]
        .split(":")[1],
      Price: textByLine[x]
        .substring(
          textByLine[x].indexOf("{") + 1,
          textByLine[x].lastIndexOf("}")
        )
        .split(";")[2]
        .split(":")[1],
    });
  }
}

// console.log(validRecord)
var excel = require('excel4node');
var workbook = new excel.Workbook();
var worksheet = workbook.addWorksheet('Sheet 1');



for (x = 0; x < validRecord.length; x++) {
  worksheet
    .cell(x + 2, 1)
    .string(
      validRecord[x]["Date"]
    );
  worksheet.cell(x + 2, 2).string(      validRecord[x]["Name"]);
  worksheet.cell(x + 2, 3).string(      validRecord[x]["Product"]);
  worksheet.cell(x + 2, 4).string(      validRecord[x]["Quantity"] );
  worksheet.cell(x + 2, 5).string(      validRecord[x]["Price"]);
}
// Set value of cell A1 to 100 as a number type styled with paramaters of style
// worksheet.cell(1,1).number(100).style(style);

// // Set value of cell B1 to 300 as a number type styled with paramaters of style
// worksheet.cell(1,2).number(200).style(style);

// // Set value of cell C1 to a formula styled with paramaters of style
// worksheet.cell(1,3).formula('A1 + B1').style(style);

// // Set value of cell A2 to 'string' styled with paramaters of style
// worksheet.cell(2,1).string('string').style(style);

// // Set value of cell A3 to true as a boolean type styled with paramaters of style but with an adjustment to the font size.
// worksheet.cell(3,1).bool(true).style(style).style({font: {size: 14}});

// workbook.write("NewExcel3.xlsx");

document.querySelectorAll(".drop-zone__input").forEach((inputElement) => {
  const dropZoneElement = inputElement.closest(".drop-zone");

  dropZoneElement.addEventListener("click", (e) => {
    inputElement.click();
  });

  inputElement.addEventListener("change", (e) => {
    if (inputElement.files.length) {
      updateThumbnail(dropZoneElement, inputElement.files[0]);
    }
  });

  dropZoneElement.addEventListener("dragover", (e) => {
    e.preventDefault();
    dropZoneElement.classList.add("drop-zone--over");
  });

  ["dragleave", "dragend"].forEach((type) => {
    dropZoneElement.addEventListener(type, (e) => {
      dropZoneElement.classList.remove("drop-zone--over");
    });
  });

  dropZoneElement.addEventListener("drop", (e) => {
    e.preventDefault();

    if (e.dataTransfer.files.length) {
      inputElement.files = e.dataTransfer.files;
      updateThumbnail(dropZoneElement, e.dataTransfer.files[0]);
    }

    dropZoneElement.classList.remove("drop-zone--over");
  });
});

/**
 * Updates the thumbnail on a drop zone element.
 *
 * @param {HTMLElement} dropZoneElement
 * @param {File} file
 */
function updateThumbnail(dropZoneElement, file) {
  let thumbnailElement = dropZoneElement.querySelector(".drop-zone__thumb");

  // First time - remove the prompt
  if (dropZoneElement.querySelector(".drop-zone__prompt")) {
    dropZoneElement.querySelector(".drop-zone__prompt").remove();
  }

  // First time - there is no thumbnail element, so lets create it
  if (!thumbnailElement) {
    thumbnailElement = document.createElement("div");
    thumbnailElement.classList.add("drop-zone__thumb");
    dropZoneElement.appendChild(thumbnailElement);
  }

  thumbnailElement.dataset.label = file.name;

  // Show thumbnail for image files
  if (file.type.startsWith("image/")) {
    const reader = new FileReader();

    reader.readAsDataURL(file);
    reader.onload = () => {
      thumbnailElement.style.backgroundImage = `url('${reader.result}')`;
    };
  } else {
    thumbnailElement.style.backgroundImage = null;
  }
}
