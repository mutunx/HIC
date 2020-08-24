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
            <Table :columns="columns" :data="infos">
              <template slot-scope="{ row, index }" slot="name">
                <Input type="text" v-model="editName" v-if="editIndex === index" />
                <span v-else>{{ row.name }}</span>
              </template>

              <template slot-scope="{ row, index }" slot="age">
                <Input type="text" v-model="editAge" v-if="editIndex === index" />
                <span v-else>{{ row.age }}</span>
              </template>

              <template slot-scope="{ row, index }" slot="birthday">
                <Input type="text" v-model="editBirthday" v-if="editIndex === index" />
                <span v-else>{{ getBirthday(row.birthday) }}</span>
              </template>

              <template slot-scope="{ row, index }" slot="address">
                <Input type="text" v-model="editAddress" v-if="editIndex === index" />
                <span v-else>{{ row.address }}</span>
              </template>

              <template slot-scope="{ row, index }" slot="action">
                <div v-if="editIndex === index">
                  <Button @click="handleSave(index)">保存</Button>
                  <Button @click="editIndex = -1">取消</Button>
                </div>
                <div v-else>
                  <Button @click="handleEdit(row, index)">操作</Button>
                </div>
              </template>
            </Table>
          </div>
        </Card>
      </Content>
      <Footer class="layout-footer-center">2011-2016 &copy; TalkingData</Footer>
    </Layout>
  </div>
</template>
<script>
export default {
  data() {
    return {
      columns: [
        {
          title: "姓名",
          slot: "name",
        },
        {
          title: "性别",
          slot: "sex",
        },
        {
          title: "民族",
          slot: "nation",
        },
        {
          title: "出生日期",
          slot: "birthday",
        },
        {
          title: "籍贯",
          slot: "hometown",
        },
        {
          title: "职业",
          slot: "career",
        },
        {
          title: "身份证",
          slot: "IDCard",
        },
        {
          title: "单位",
          slot: "unit",
        },
        {
          title: "现居地",
          slot: "living",
        },
        {
          title: "父亲",
          slot: "father",
        },
        {
          title: "母亲",
          slot: "mother",
        },
        {
          title: "教育背景",
          slot: "education",
        },
        {
          title: "操作",
          slot: "action",
        },
      ],
      infos: [],
      editIndex: -1, // 当前聚焦的输入框的行数
      editName: "", // 第一列输入框，当然聚焦的输入框的输入内容，与 data 分离避免重构的闪烁
      editAge: "", // 第二列输入框
      editBirthday: "", // 第三列输入框
      editAddress: "", // 第四列输入框
    };
  },
  created() {
   var that=this
    this.$axios.get('infos').then(function(res){
              console.log(res.data.data.data)
                that.infos =   res.data.data.data 
            });
  },
  methods: {
      getData() {
            this.$axios.get('infos').then(function(res){
              console.log(res.data.data.data)
                return  res.data.data.data 
            });

      },
    handleEdit(row, index) {
      this.editName = row.name;
      this.editAge = row.age;
      this.editAddress = row.address;
      this.editBirthday = row.birthday;
      this.editIndex = index;
    },
    handleSave(index) {
      this.data[index].name = this.editName;
      this.data[index].age = this.editAge;
      this.data[index].birthday = this.editBirthday;
      this.data[index].address = this.editAddress;
      this.editIndex = -1;
    },
    getBirthday(birthday) {
      const date = new Date(parseInt(birthday));
      const year = date.getFullYear();
      const month = date.getMonth() + 1;
      const day = date.getDate();
      return `${year}-${month}-${day}`;
    },
  },
};
</script>
