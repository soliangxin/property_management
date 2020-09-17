<template>
  <div class="app-container">
    <el-table
      v-loading="listLoading"
      :data="list"
      element-loading-text="Loading"
      border
      fit
      highlight-current-row
    >
      <el-table-column align="center" label="收费项目编号" width="160">
        <template slot-scope="scope">
          {{ scope.row.code }}
        </template>
      </el-table-column>
      <el-table-column label="收费项目名称" width="190" align="center">
        <template slot-scope="scope">
          {{ scope.row.name }}
        </template>
      </el-table-column>
      <el-table-column label="所属小区编号" width="190" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.community_code }}</span>
        </template>
      </el-table-column>
      <el-table-column align="center" prop="created_at" label="添加时间" width="210">
        <template slot-scope="scope">
          <i class="el-icon-time" />
          <span>{{ scope.row.create_time }}</span>
        </template>
      </el-table-column>
      <el-table-column align="center" prop="created_at" label="修改时间" width="210">
        <template slot-scope="scope">
          <i class="el-icon-time" />
          <span>{{ scope.row.update_time }}</span>
        </template>
      </el-table-column>
      <el-table-column label="备注">
        <template slot-scope="scope">
          {{ scope.row.developer }}
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script>
import { getChangeItemList } from '@/api/charge'

export default {
  filters: {
    statusFilter(status) {
      const statusMap = {
        published: 'success',
        draft: 'gray',
        deleted: 'danger'
      }
      return statusMap[status]
    }
  },
  data() {
    return {
      list: null,
      listLoading: true
    }
  },
  created() {
    this.fetchData()
  },
  methods: {
    fetchData() {
      this.listLoading = true
      getChangeItemList().then(response => {
        this.list = response.data.items
        this.listLoading = false
      })
    }
  }
}
</script>
