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
      <el-table-column align="center" label="所属小区编号" width="190">
        <template slot-scope="scope">
          {{ scope.row.community_code }}
        </template>
      </el-table-column>
      <el-table-column align="center" label="值班人名称" width="210">
        <template slot-scope="scope">
          {{ scope.row.name }}
        </template>
      </el-table-column>
      <el-table-column align="center" prop="created_at" label="值班开始时间" :formatter="dateFormat" width="230">
        <template slot-scope="scope">
          <i class="el-icon-time" />
          <span>{{ scope.row.start_time }}</span>
        </template>
      </el-table-column>
      <el-table-column align="center" prop="created_at" label="值班结束时间" :formatter="dateFormat" width="230">
        <template slot-scope="scope">
          <i class="el-icon-time" />
          <span>{{ scope.row.end_time }}</span>
        </template>
      </el-table-column>
      <el-table-column label="备注">
        <template slot-scope="scope">
          {{ scope.row.remark }}
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script>
import { getWatchList } from '@/api/watch'

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
      getWatchList().then(response => {
        this.list = response.data.items
        this.listLoading = false
      })
    }
  }
}
</script>
