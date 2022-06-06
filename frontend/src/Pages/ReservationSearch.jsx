import { useState } from "react";
import InputLabel from "@mui/material/InputLabel";
import MenuItem from "@mui/material/MenuItem";
import FormHelperText from "@mui/material/FormHelperText";
import FormControl from "@mui/material/FormControl";
import Select from "@mui/material/Select";
import TextField from "@mui/material/TextField";
import Table from "@mui/material/Table";
import TableBody from "@mui/material/TableBody";
import TableCell from "@mui/material/TableCell";
import TableContainer from "@mui/material/TableContainer";
import TableHead from "@mui/material/TableHead";
import TableRow from "@mui/material/TableRow";
import Box from "@mui/material/Box";
import Button from "@mui/material/Button";
import Typography from "@mui/material/Typography";
import Paper from "@mui/material/Paper";
import { HospitalAPI, ReservationAPI } from "../api";

export default function ReservationSearch() {
  const area_data = {
    臺北市: [
      "中正區",
      "大同區",
      "中山區",
      "萬華區",
      "信義區",
      "松山區",
      "大安區",
      "南港區",
      "北投區",
      "內湖區",
      "士林區",
      "文山區",
    ],
    新北市: [
      "板橋區",
      "新莊區",
      "泰山區",
      "林口區",
      "淡水區",
      "金山區",
      "八里區",
      "萬里區",
      "石門區",
      "三芝區",
      "瑞芳區",
      "汐止區",
      "平溪區",
      "貢寮區",
      "雙溪區",
      "深坑區",
      "石碇區",
      "新店區",
      "坪林區",
      "烏來區",
      "中和區",
      "永和區",
      "土城區",
      "三峽區",
      "樹林區",
      "鶯歌區",
      "三重區",
      "蘆洲區",
      "五股區",
    ],
    基隆市: [
      "仁愛區",
      "中正區",
      "信義區",
      "中山區",
      "安樂區",
      "暖暖區",
      "七堵區",
    ],
    桃園市: [
      "桃園區",
      "中壢區",
      "平鎮區",
      "八德區",
      "楊梅區",
      "蘆竹區",
      "龜山區",
      "龍潭區",
      "大溪區",
      "大園區",
      "觀音區",
      "新屋區",
      "復興區",
    ],
    新竹縣: [
      "竹北市",
      "竹東鎮",
      "新埔鎮",
      "關西鎮",
      "峨眉鄉",
      "寶山鄉",
      "北埔鄉",
      "橫山鄉",
      "芎林鄉",
      "湖口鄉",
      "新豐鄉",
      "尖石鄉",
      "五峰鄉",
    ],
    新竹市: ["東區", "北區", "香山區"],
    苗栗縣: [
      "苗栗市",
      "通霄鎮",
      "苑裡鎮",
      "竹南鎮",
      "頭份鎮",
      "後龍鎮",
      "卓蘭鎮",
      "西湖鄉",
      "頭屋鄉",
      "公館鄉",
      "銅鑼鄉",
      "三義鄉",
      "造橋鄉",
      "三灣鄉",
      "南庄鄉",
      "大湖鄉",
      "獅潭鄉",
      "泰安鄉",
    ],
    臺中市: [
      "中區",
      "東區",
      "南區",
      "西區",
      "北區",
      "北屯區",
      "西屯區",
      "南屯區",
      "太平區",
      "大里區",
      "霧峰區",
      "烏日區",
      "豐原區",
      "后里區",
      "東勢區",
      "石岡區",
      "新社區",
      "和平區",
      "神岡區",
      "潭子區",
      "大雅區",
      "大肚區",
      "龍井區",
      "沙鹿區",
      "梧棲區",
      "清水區",
      "大甲區",
      "外埔區",
      "大安區",
    ],
    南投縣: [
      "南投市",
      "埔里鎮",
      "草屯鎮",
      "竹山鎮",
      "集集鎮",
      "名間鄉",
      "鹿谷鄉",
      "中寮鄉",
      "魚池鄉",
      "國姓鄉",
      "水里鄉",
      "信義鄉",
      "仁愛鄉",
    ],
    彰化縣: [
      "彰化市",
      "員林鎮",
      "和美鎮",
      "鹿港鎮",
      "溪湖鎮",
      "二林鎮",
      "田中鎮",
      "北斗鎮",
      "花壇鄉",
      "芬園鄉",
      "大村鄉",
      "永靖鄉",
      "伸港鄉",
      "線西鄉",
      "福興鄉",
      "秀水鄉",
      "埔心鄉",
      "埔鹽鄉",
      "大城鄉",
      "芳苑鄉",
      "竹塘鄉",
      "社頭鄉",
      "二水鄉",
      "田尾鄉",
      "埤頭鄉",
      "溪州鄉",
    ],
    雲林縣: [
      "斗六市",
      "斗南鎮",
      "虎尾鎮",
      "西螺鎮",
      "土庫鎮",
      "北港鎮",
      "莿桐鄉",
      "林內鄉",
      "古坑鄉",
      "大埤鄉",
      "崙背鄉",
      "二崙鄉",
      "麥寮鄉",
      "臺西鄉",
      "東勢鄉",
      "褒忠鄉",
      "四湖鄉",
      "口湖鄉",
      "水林鄉",
      "元長鄉",
    ],
    嘉義縣: [
      "太保市",
      "朴子市",
      "布袋鎮",
      "大林鎮",
      "民雄鄉",
      "溪口鄉",
      "新港鄉",
      "六腳鄉",
      "東石鄉",
      "義竹鄉",
      "鹿草鄉",
      "水上鄉",
      "中埔鄉",
      "竹崎鄉",
      "梅山鄉",
      "番路鄉",
      "大埔鄉",
      "阿里山鄉",
    ],
    嘉義市: ["東區", "西區"],
    臺南市: [
      "中西區",
      "東區",
      "南區",
      "北區",
      "安平區",
      "安南區",
      "永康區",
      "歸仁區",
      "新化區",
      "左鎮區",
      "玉井區",
      "楠西區",
      "南化區",
      "仁德區",
      "關廟區",
      "龍崎區",
      "官田區",
      "麻豆區",
      "佳里區",
      "西港區",
      "七股區",
      "將軍區",
      "學甲區",
      "北門區",
      "新營區",
      "後壁區",
      "白河區",
      "東山區",
      "六甲區",
      "下營區",
      "柳營區",
      "鹽水區",
      "善化區",
      "大內區",
      "山上區",
      "新市區",
      "安定區",
    ],
    高雄市: [
      "楠梓區",
      "左營區",
      "鼓山區",
      "三民區",
      "鹽埕區",
      "前金區",
      "新興區",
      "苓雅區",
      "前鎮區",
      "小港區",
      "旗津區",
      "鳳山區",
      "大寮區",
      "鳥松區",
      "林園區",
      "仁武區",
      "大樹區",
      "大社區",
      "岡山區",
      "路竹區",
      "橋頭區",
      "梓官區",
      "彌陀區",
      "永安區",
      "燕巢區",
      "田寮區",
      "阿蓮區",
      "茄萣區",
      "湖內區",
      "旗山區",
      "美濃區",
      "內門區",
      "杉林區",
      "甲仙區",
      "六龜區",
      "茂林區",
      "桃源區",
      "那瑪夏區",
    ],
    屏東縣: [
      "屏東市",
      "潮州鎮",
      "東港鎮",
      "恆春鎮",
      "萬丹鄉",
      "長治鄉",
      "麟洛鄉",
      "九如鄉",
      "里港鄉",
      "鹽埔鄉",
      "高樹鄉",
      "萬巒鄉",
      "內埔鄉",
      "竹田鄉",
      "新埤鄉",
      "枋寮鄉",
      "新園鄉",
      "崁頂鄉",
      "林邊鄉",
      "南州鄉",
      "佳冬鄉",
      "琉球鄉",
      "車城鄉",
      "滿州鄉",
      "枋山鄉",
      "霧台鄉",
      "瑪家鄉",
      "泰武鄉",
      "來義鄉",
      "春日鄉",
      "獅子鄉",
      "牡丹鄉",
      "三地門鄉",
    ],
    宜蘭縣: [
      "宜蘭市",
      "羅東鎮",
      "蘇澳鎮",
      "頭城鎮",
      "礁溪鄉",
      "壯圍鄉",
      "員山鄉",
      "冬山鄉",
      "五結鄉",
      "三星鄉",
      "大同鄉",
      "南澳鄉",
    ],
    花蓮縣: [
      "花蓮市",
      "鳳林鎮",
      "玉里鎮",
      "新城鄉",
      "吉安鄉",
      "壽豐鄉",
      "秀林鄉",
      "光復鄉",
      "豐濱鄉",
      "瑞穗鄉",
      "萬榮鄉",
      "富里鄉",
      "卓溪鄉",
    ],
    臺東縣: [
      "臺東市",
      "成功鎮",
      "關山鎮",
      "長濱鄉",
      "海端鄉",
      "池上鄉",
      "東河鄉",
      "鹿野鄉",
      "延平鄉",
      "卑南鄉",
      "金峰鄉",
      "大武鄉",
      "達仁鄉",
      "綠島鄉",
      "蘭嶼鄉",
      "太麻里鄉",
    ],
    澎湖縣: ["馬公市", "湖西鄉", "白沙鄉", "西嶼鄉", "望安鄉", "七美鄉"],
    金門縣: ["金城鎮", "金湖鎮", "金沙鎮", "金寧鄉", "烈嶼鄉", "烏坵鄉"],
    連江縣: ["南竿鄉", "北竿鄉", "莒光鄉", "東引鄉"],
  };
  const counties = Object.keys(area_data);
  const [county, setCounty] = useState(counties[0]);
  const handleChange = (event) => {
    setCounty(event.target.value);
  };
  const handleSearch = () => {
    HospitalAPI.getHospital();
  };
  return (
    <>
      <Box
        sx={{
          display: "flex",
          justifyContent: "center",
          alignContent: "center",
          alignItems: "center",
          mb: 1,
        }}
      >
        <Typography>縣市：</Typography>
        <FormControl sx={{ mr: 1 }}>
          <Select
            value={county}
            onChange={handleChange}
            displayEmpty
            sx={{ mx: 1, height: 40 }}
          >
            {counties.map((ele) => (
              <MenuItem
                key={ele}
                value={ele}
                sx={{
                  fontSize: 8,
                }}
              >
                {ele}
              </MenuItem>
            ))}
          </Select>
        </FormControl>
        <Button
          variant="contained"
          sx={{
            mx: 1,
          }}
          onClick={handleSearch}
        >
          查詢
        </Button>
      </Box>
      {area_data[county].map((ele) => {
        return (
          <Box>
            <Typography>{ele}</Typography>
            <TableContainer component={Paper}>
              <Table sx={{ minWidth: 650 }} aria-label="simple table">
                <TableHead>
                  <TableRow sx={{ height: 20 }}>
                    <TableCell>醫療院所名稱</TableCell>
                    <TableCell align="right">疫苗種類</TableCell>
                    <TableCell align="right">剩餘劑數</TableCell>
                  </TableRow>
                </TableHead>
                <TableBody></TableBody>
              </Table>
            </TableContainer>
          </Box>
        );
      })}
    </>
  );
}
