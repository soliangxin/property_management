<template>
  <div class="app-container">
    <el-form ref="form" :model="form" label-width="120px">
      <el-form-item label="房产编号">
        <el-input v-model="form.name" placeholder="建议CM开头" />
      </el-form-item>
      <el-form-item label="房产名称">
        <el-input v-model="form.name2" />
      </el-form-item>
      <el-form-item label="户主名称">
        <el-input v-model="form.name3" />
      </el-form-item>
      <el-form-item label="户主联系方式">
        <el-input v-model="form.name4" />
      </el-form-item>
      <el-form-item label="房间数">
        <el-input v-model="form.name5" />
      </el-form-item>
      <el-form-item label="单元信息">
        <el-input v-model="form.name6" />
      </el-form-item>
      <el-form-item label="楼层信息">
        <el-input v-model="form.name7" />
      </el-form-item>
      <el-form-item label="所属小区编号">
        <el-input v-model="form.name8" />
      </el-form-item>
      <el-form-item label="栋数编号">
        <el-input v-model="form.name9" />
      </el-form-item>
      <el-form-item label="房产描述">
        <el-input v-model="form.desc" type="textarea" />
      </el-form-item>
      <el-form-item label="入住时间">
        <el-date-picker v-model="form.date" type="date" style="width: 100%;" />
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
      this.$message('submit!')
    },
    onCancel() {
      this.$message({
        message: 'cancel!',
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
