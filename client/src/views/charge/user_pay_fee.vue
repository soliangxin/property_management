<template>
  <div class="app-container">
    <el-form ref="form" :model="form" label-width="120px">
      <el-form-item label="房产编号">
        <el-input v-model="form.code" />
      </el-form-item>
      <el-form-item label="缴费项目编号">
        <el-input v-model="form.name" type="name" />
      </el-form-item>
      <el-form-item label="应收金额">
        <el-input v-model="form.name" type="name" />
      </el-form-item>
      <el-form-item label="实收金额">
        <el-input v-model="form.name" type="name" />
      </el-form-item>
      <el-form-item label="所属小区编号">
        <el-input v-model="form.community_code" type="name" />
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
        code: '',
        name: '',
        status: '',
        community_code: '',
        add_date: '',
        add_time: ''
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
    },
    created() {
    },
    mounted() {
    }
  }
}
</script>

<style scoped>
  .line{
    text-align: center;
  }
</style>
