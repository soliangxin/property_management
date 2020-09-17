<template>
  <div class="app-container">
    <el-form ref="form" :model="form" label-width="120px">
      <el-form-item label="房产编号">
        <el-input v-model="form.name" placeholder="建议HS开头" />
      </el-form-item>
      <el-form-item label="小区名称">
        <el-input v-model="form.name1" />
      </el-form-item>
      <el-form-item label="简介">
        <el-input v-model="form.desc" type="textarea" />
      </el-form-item>
      <el-form-item label="坐落地址">
        <el-input v-model="form.desc1" type="textarea" />
      </el-form-item>
      <el-form-item label="占地面积">
        <el-input v-model="form.name2" placeholder="单位：平米" />
      </el-form-item>
      <el-form-item label="开发商名称">
        <el-input v-model="form.name3" />
      </el-form-item>
      <el-form-item label="物业公司名称">
        <el-input v-model="form.name4" />
      </el-form-item>
      <el-form-item label="绿化率">
        <el-input v-model="form.name5" placeholder="单位：百分比" />
      </el-form-item>
      <el-form-item label="总栋数">
        <el-input v-model="form.name6" />
      </el-form-item>
      <el-form-item label="总户数">
        <el-input v-model="form.name7" />
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
