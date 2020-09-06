<style scoped>
.layout {
  border: 1px solid #d7dde4;
  background: #f5f7f9;
  position: relative;
  border-radius: 4px;
  overflow: hidden;
}
.layout-logo {
  width: 100px;
  height: 30px;
  background: #5b6270;
  border-radius: 3px;
  float: left;
  position: relative;
  top: 15px;
  left: 20px;
}
.layout-nav {
  width: 420px;
  margin: 0 auto;
  margin-right: 20px;
}
.layout-footer-center {
  text-align: center;
}
</style>
<template>
  <div class="layout">
    <Layout>
      <Header>
        <Menu mode="horizontal" theme="dark" active-name="1">
          <div class="layout-logo"></div>
          <div class="layout-nav">
            <MenuItem name="1">
              <Icon type="ios-navigate"></Icon>获取信息
            </MenuItem>
            <MenuItem name="2">
              <Icon type="ios-keypad"></Icon>Item 2
            </MenuItem>
            <MenuItem name="3">
              <Icon type="ios-analytics"></Icon>Item 3
            </MenuItem>
            <MenuItem name="4">
              <Icon type="ios-paper"></Icon>Item 4
            </MenuItem>
          </div>
        </Menu>
      </Header>
      <Content :style="{padding: '0 50px'}">
        <Breadcrumb :style="{margin: '20px 0'}">
          <BreadcrumbItem>Home</BreadcrumbItem>
          <BreadcrumbItem>Components</BreadcrumbItem>
          <BreadcrumbItem>Layout</BreadcrumbItem>
        </Breadcrumb>
        <Card>
          <div style="min-height: 200px;">
            <Row>
              <Col span="6">
                <Cascader :data="areas" v-model="searchArea" ></Cascader>
                </Col>
              <Col span="6">
                <DatePicker type="daterange" v-model="searchTime" placement="bottom-end" placeholder="Select month" ></DatePicker>
                </Col>
              <Col span="6">
               <Input suffix="ios-search" v-model="searchWord" placeholder="Enter text" style="width: auto" />
                </Col>
                
              <Col span="6">
               <Upload action="/hic/import/info">
               <Button :size="buttonSize" icon="ios-cloud-upload" style="text-align:end"  type="primary">导入</Button>
               </Upload>
                </Col>
            </Row>
             
             
            <Table :columns="columns" :data="infos">

            </Table>
          </div>
        </Card>
      </Content>
      <Footer class="layout-footer-center">2011-2016 &copy; TalkingData</Footer>
    </Layout>
  </div>
</template>
<script>
import moment from 'moment'
export default {
  data() {
    return {
      buttonSize:'large',
      columns: [
        {
          title: "姓名",
          key: "Name",
        },
        {
          title: "性别",
          key: "Sex",
        },
        {
          title: "省份",
          key: "Province",
        },
        {
          title: "出生日期",
          key: "Birthday",
        },
        {
          title: "市县",
          key: "City",
        },
        {
          title: "地区",
          key: "Area",
        },
        {
          title: "身份证",
          key: "ID",
        },
        {
          title: "验证通过",
          key: "ValidState",
        },

      ],
      searchWord:"",
      searchTime:[],
      searchArea:[],
      infos: [],
      areas:[],
      editIndex: -1, // 当前聚焦的输入框的行数
      editName: "", // 第一列输入框，当然聚焦的输入框的输入内容，与 data 分离避免重构的闪烁
      editBirthday: "", // 第三列输入框
    };
  },
  watch:{
    
    searchWord: function (newWord, oldWord) {
      if (newWord === "") {
        return 
      }
      var that=this
      this.$axios.get('search/info?keyword='+newWord).then(function(res){
        console.log(res.data.data)
                that.infos =   res.data.data
            });
    },
    
    searchTime: function (newInput, oldInput) {
      if (newInput[0] === "") {
        return 
      }
      console.log(moment(newInput[0]).format("YYYY-MM-DD"),oldInput)
      let sTime = moment(newInput[0]).format("YYYY-MM-DD")
      let eTime = moment(newInput[1]).format("YYYY-MM-DD")
      var that=this
      this.$axios.get('search/info?sTime='+sTime+"&eTime="+eTime).then(function(res){
        console.log(res.data.data)
                that.infos =   res.data.data
            });
    },
    
    searchArea: function (newInput) {
      console.log(newInput)
      if (newInput === "") {
        return 
      }
      var that=this
      this.$axios.get('search/info?keyword='+newInput[newInput.length-1]).then(function(res){
        console.log(res.data.data)
                that.infos =   res.data.data
            });
    }
  },
  created() {
   var that=this
    this.$axios.get('infos').then(function(res){
                that.infos =   res.data.data.data 
            });
            
    this.$axios.get('adc').then(function(res){
          let data = res.data.data
          console.log(data)
                that.areas =   data
            });
  },
  methods: {
     loadData (item, callback) {
              
            },
      getData() {
            this.$axios.get('infos').then(function(res){
              console.log(res.data.data.data)
                return  res.data.data.data 
            });

      },
    handleEdit(row, index) {
      this.editName = row.name;
      this.editBirthday = row.birthday;
      this.editIndex = index;
    },
    handleSave(index) {
      this.data[index].name = this.editName;
      this.data[index].birthday = this.editBirthday;
      this.editIndex = -1;
    },
  }
};
</script>
