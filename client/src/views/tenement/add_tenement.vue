<template>
  <div class="app-container">
    <el-form ref="form" :model="form" label-width="120px">
      <el-form-item label="房产编号">
        <el-input v-model="form.name" placeholder="建议HS开头" />
      </el-form-item>
      <el-form-item label="姓名">
        <el-input v-model="form.name1" />
      </el-form-item>
      <el-form-item label="身份证号">
        <el-input v-model="form.name2" />
      </el-form-item>
      <el-form-item label="联系方式">
        <el-input v-model="form.name3" />
      </el-form-item>
      <el-form-item label="职业">
        <el-input v-model="form.name4" />
      </el-form-item>
      <el-form-item label="出生日期">
        <el-date-picker v-model="form.date" type="date" style="width: 100%;" />
      </el-form-item>
      <el-form-item label="性别">
        <el-select v-model="form.region" placeholder="请选择性别">
          <el-option label="男" value="0" />
          <el-option label="女" value="1" />
        </el-select>
      </el-form-item>
      <el-form-item label="成员类型">
        <el-select v-model="form.region1" placeholder="成员类型">
          <el-option label="户主" value="0" />
          <el-option label="家庭成员" value="1" />
          <el-option label="租户" value="3" />
        </el-select>
      </el-form-item>
      <el-form-item label="所属小区编号">
        <el-input v-model="form.name6" />
      </el-form-item>
      <el-form-item label="入住日期">
        <el-date-picker v-model="form.date1" type="date" style="width: 100%;" />
      </el-form-item>
      <el-form-item label="备注">
        <el-input v-model="form.desc1" type="textarea" />
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="onSubmit">提交</el-button>
        <el-button @click="onCancel">取消</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
export default {
  data() {
    return {
      form: {
        name: '',
        region: '',
        date1: '',
        date2: '',
        delivery: false,
        type: [],
        resource: '',
        desc: ''
      }
    }
  },
  methods: {
    onSubmit() {
      this.$message('已提交')
    },
    onCancel() {
      this.$message({
        message: '已取消',
        type: 'warning'
      })
    }
  }
}
</script>

<style scoped>
  .line{
    text-align: center;
  }
</style>
