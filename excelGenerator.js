// async function loadLocaltxtfile() {
//   const res = await fetch("/getTxtfile"); // DEFAULT 係Get Method
//   const stocks = await res.json();

//   for (let stock of stocks) {
//     NonExpiredItems.push([stock.name, stock.quantity]);
//   }
//   // const reducer = (accumulator, currentValue) => parseInt(accumulator) + parseInt(currentValue);
//   // document.querySelector("#Inventory").innerHTML = `Remaining Stock of ${product}: ${NonExpiredItems.reduce(reducer)}`;
//   // document.querySelector("#Inventory").style="color:black;text-align:center;";
//   console.log(NonExpiredItems);
// }
// window.onload = (event) => {
//   loadLocaltxtfile();
// };
let mybutton = document.getElementById("send");
 
mybutton.addEventListener("submit", (event) => {

    event.preventDefault();
    const form = event.target;
    // 1. JSON Object
    // const formObj = {
    //     content: form.content.value
    // };
    // const res = await fetch('/memos',{
    //     method:"POST",
    //     headers:{
    //         "Content-Type":"application/json"
    //     },
    //     body: JSON.stringify(formObj)
    // });
    const formData = new FormData();
    formData.append("contactNumber", form.contactNumber.value);

    let obj = {
      fileInput :form.fileInput.value,
    };

    console.log(obj)

    const res = fetch("/newFileInput", {
      method: "POST",
      headers: {
        "Content-Type": "application/json; charset=utf-8",
      },
      body: JSON.stringify(obj)
    });

      let result = res.json();

      console.log(result)
      console.log(result.fileInput)

    
    if (res.status == 200) {
      // guaranteed 返個Response 真係啱嘅response
      console.log(result);
    }

  

});
