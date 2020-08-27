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
              <Col span="12">
                <Cascader :data="areas" change-on-select></Cascader>
                </Col>
              <Col span="12">
                <DatePicker type="daterange" placement="bottom-end" placeholder="Select date" style="width: 200px"></DatePicker>
                </Col>
              <Col span="12">
               <Input suffix="ios-search" placeholder="Enter text" style="width: auto" />
                </Col>
                
              <Col span="12">
               <Button :size="buttonSize" icon="ios-cloud-upload"  type="primary">导入</Button>
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
export default {
  data() {
    return {
      columns: [
        {
          title: "姓名",
          key: "name",
        },
        {
          title: "性别",
          key: "sex",
        },
        {
          title: "民族",
          key: "nation",
        },
        {
          title: "出生日期",
          key: "birthday",
        },
        {
          title: "籍贯",
          key: "hometown",
        },
        {
          title: "职业",
          key: "career",
        },
        {
          title: "身份证",
          key: "IDCard",
        },
        {
          title: "单位",
          key: "unit",
        },
        {
          title: "现居地",
          key: "living",
        },
        {
          title: "父亲",
          key: "father",
        },
        {
          title: "母亲",
          key: "mother",
        },
        {
          title: "教育背景",
          key: "education",
        },
        {
          title: "操作",
          key: "action",
        },
      ],
      infos: [],
      areas:[{
                    value: 'beijing',
                    label: '北京',
                    children: [
                        {
                            value: 'gugong',
                            label: '故宫'
                        },
                        {
                            value: 'tiantan',
                            label: '天坛'
                        },
                        {
                            value: 'wangfujing',
                            label: '王府井'
                        }
                    ]
                }, {
                    value: 'jiangsu',
                    label: '江苏',
                    children: [
                        {
                            value: 'nanjing',
                            label: '南京',
                            children: [
                                {
                                    value: 'fuzimiao',
                                    label: '夫子庙',
                                }
                            ]
                        },
                        {
                            value: 'suzhou',
                            label: '苏州',
                            children: [
                                {
                                    value: 'zhuozhengyuan',
                                    label: '拙政园',
                                },
                                {
                                    value: 'shizilin',
                                    label: '狮子林',
                                }
                            ]
                        }
                    ],
                }],
      editIndex: -1, // 当前聚焦的输入框的行数
      editName: "", // 第一列输入框，当然聚焦的输入框的输入内容，与 data 分离避免重构的闪烁
      editBirthday: "", // 第三列输入框
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
      this.editBirthday = row.birthday;
      this.editIndex = index;
    },
    handleSave(index) {
      this.data[index].name = this.editName;
      this.data[index].birthday = this.editBirthday;
      this.editIndex = -1;
    },

  },
};
</script>
