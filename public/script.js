const ckinp = document.getElementById("ck-inp")
const ckbtn = document.getElementById("ck-btn")
const ckmsg = document.getElementById("ck-msg")
const code = document.getElementById("code")
let pass = true 

async function check_pass(){
  const res = await fetch("/api/password");
  const json = await res.json() 
  pass = json["enabled"]
  if (!pass){
    ckinp.style = "display: none;"
  }
}

ckbtn.addEventListener("click", async function(){
  let res
  if (pass)
    res = await fetch("/api/gen?password="+ckinp.value)
  else
    res = await fetch("/api/gen?password=")

  const json = await res.json()
  if (json["error"] != ""){
    ckmsg.innerText = json["error"]
    return ckmsg.style = "color: red;"
  }
  
  ckmsg.innerText = json["key"]
  return ckmsg.style = "color: white;"
})

code.innerText = "passctl server "+document.URL
check_pass()
