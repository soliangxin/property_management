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
      <el-table-column align="center" label="房产编号" width="150">
        <template slot-scope="scope">
          {{ scope.row.house_code }}
        </template>
      </el-table-column>
      <el-table-column align="center" label="房产名称" width="120">
        <template slot-scope="scope">
          {{ scope.row.name }}
        </template>
      </el-table-column>
      <el-table-column label="户主" width="110" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.owner_name }}</span>
        </template>
      </el-table-column>
      <el-table-column label="房间数" width="80" align="center">
        <template slot-scope="scope">
          {{ scope.row.rooms }}
        </template>
      </el-table-column>
      <el-table-column label="所属小区编号" width="170" align="center">
        <template slot-scope="scope">
          {{ scope.row.community_code }}
        </template>
      </el-table-column>
      <el-table-column label="楼层信息" width="90" align="center">
        <template slot-scope="scope">
          {{ scope.row.floor }}
        </template>
      </el-table-column>
      <el-table-column label="单元信息" width="100" align="center">
        <template slot-scope="scope">
          {{ scope.row.unit }}
        </template>
      </el-table-column>
      <el-table-column label="户主联系方式" width="120" align="center">
        <template slot-scope="scope">
          {{ scope.row.owner_tel }}
        </template>
      </el-table-column>
      <el-table-column label="房产描述">
        <template slot-scope="scope">
          {{ scope.row.house_desc }}
        </template>
      </el-table-column>
      <el-table-column align="center" prop="created_at" label="入住时间" :formatter="dateFormat" width="200">
        <template slot-scope="scope">
          <i class="el-icon-time" />
          <span>{{ scope.row.enter_time }}</span>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script>
import { getHouseList } from '@/api/house'

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
      getHouseList().then(response => {
        this.list = response.data.items
        this.listLoading = false
      })
    }
  }
}
</script>
