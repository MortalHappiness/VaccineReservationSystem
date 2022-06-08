const FakeDataGenerator = require("fake-data-generator-taiwan");
const moment = require("moment");
// 產生一百組假資料
if (process.argv.length != 3) {
  console.error("Usage: node generate.js <user count>");
  return;
}
const NUM = parseInt(process.argv[2], 10);
let generator = new FakeDataGenerator();
function randomDate(start, end) {
  return new Date(
    start.getTime() + Math.random() * (end.getTime() - start.getTime())
  );
}
function getRandGender() {
  return Math.floor(Math.random() * 2) == 1 ? "Female" : "Male";
}

const healthCardID = "000000000000";

console.log("ID,name,healthCardID,gender,birthDay,address,phone,vaccines");

for (let i = 0; i < NUM; i++) {
  let name = generator.Name.generate();
  let phone = generator.Mobile.generate(0, 10);
  let id = generator.IDNumber.generate();
  let address = generator.Address.generate();
  let birthDay = moment(randomDate(new Date(1930, 0, 1), new Date())).format(
    "YYYY/MM/DD"
  );
  let gender = getRandGender();

  console.log(
    `${id},${name},${healthCardID},${gender},${birthDay},${address},${phone},`
  );
}
