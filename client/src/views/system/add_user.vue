<template>
  <div class="app-container">
    <el-form ref="form" :model="form" label-width="120px">
      <el-form-item label="用户名">
        <el-input v-model="form.another_name" type="name" />
      </el-form-item>
      <el-form-item label="昵称">
        <el-input v-model="form.username" type="name" />
      </el-form-item>
      <el-form-item label="密码">
        <el-input v-model="form.password" type="password" />
      </el-form-item>
      <el-form-item label="电子邮箱">
        <el-input v-model="form.email" type="email" />
      </el-form-item>
      <el-form-item label="角色">
        <el-select v-model="form.status" type="region" placeholder="请选择角色">
          <el-option label="普通员工" value="1" />
          <el-option label="小区管理员" value="2" />
          <el-option label="超级管理员" value="3" />
        </el-select>
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
        name1: '',
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
