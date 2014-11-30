####
## API Calls
####

apiCall = (method, path, params, cb) ->
  $.ajax
    url: path
    type: method
    data: params
    #data: JSON.stringify(params)
    #contentType: 'application/json'
    cache: false
    error: (xhr, status, data) ->
      console.log xhr.responseText
      cb($.parseJSON(xhr.responseText))
    success: (data) ->
      cb(data)

getInfo = (cb) ->
  apiCall 'GET', '/api/info', {}, cb

getTables = (cb) ->
  apiCall 'GET', '/api/table', {}, cb

getTableInfo = (table, cb) ->
  apiCall 'GET', '/api/table/' + table + '/info', {}, cb

getTable = (table, cb) ->
  apiCall 'GET', '/api/table/' + table, {}, cb

getTableSql = (table, cb) ->
  apiCall 'GET', '/api/table/' + table + '/sql', {}, cb

getTableIndexes = (table, cb) ->
  apiCall 'GET', '/api/table/' + table + '/indexes', {}, cb

getTableContent = (table, cb) ->
  query = 'SELECT * FROM ' + table + ';'
  executeQuery query, cb

getQuery = (query, cb) ->
  executeQuery query, cb

executeQuery = (query, cb) ->
  data =
    query: query
  apiCall 'POST', '/api/query', data, cb


####
## Table construction functions
####

buildTableStructure = (name, cb) ->
  getTableSql name, (data) ->
    $('#structure_sql').text data.sql

    getTable name, (columns) ->
      $('#table_columns tbody').empty()

      columns.forEach (item) ->
        column = '<tr>'
        column += '<th>' + item.name + '</th>'
        column += '<th>' + item.type + '</th>'
        column += '<th>' + (if item.pk then 'True' else 'False') + '</th>'
        column += '<th>' + (if item.notnull then 'True' else 'False')  + '</th>'
        def_val = if item.dflt_value == null then 'Null' else item.dflt_value
        column += '<th>' + def_val + '</th>'
        column += '</tr>'

        $('#table_columns tbody').append column

      getTableIndexes name, (columns) ->
        $('#table_indexes tbody').empty()

        if not columns? or columns.length < 1
          return cb()

        columns.forEach (item) ->
          column = '<tr>'
          column += '<th>' + item.name + '</th>'

          cols = getColumnsFromIndexSql item.tbl_name, item.sql
          column += '<th>' + cols.join(', ') + '</th>'

          unique = if item.sql.indexOf('UNIQUE') > -1 then 'True' else 'False'
          column += '<th>' + unique + '</th>'

          sqlLink = '<a class="view-sql" '
          sqlLink += 'data-toggle="modal" data-target="#index_sql_modal" '
          sqlLink += 'data-name="' + item.name + '" '
          sqlLink += 'href="#">SQL</a>'
          pre = '<pre style="display: none;">' + item.sql + '</pre>'
          column += '<th>' + sqlLink + pre + '</th>'
          column += '</tr>'

          $('#table_indexes tbody').append column

        cb()

buildTableContent = (name, cb) ->
  getTableContent name, (data) ->
    resetResultTable()
    addHeadersToResultTable (buildResultHeader name for name in data.columns)
    addRowsToResultTable (buildResultRow row for row in data.rows)
    cb()

buildTableQueryResult = (query) ->
  getQuery query, (data) ->
    resetResultTable()
    addHeadersToResultTable (buildResultHeader name for name in data.columns)
    addRowsToResultTable (buildResultRow row for row in data.rows)


####
## Main UI building functions
####

loadTables = (cb) ->
  $('#tables li').remove

  getTables (data) ->
    data.tables.forEach (item) ->
      $('<li><span>' +  item + '</span></li>').appendTo('#tables')

    # Select the first table
    if data.tables.length > 0
      $('#tables li.selected').removeClass 'selected'
      table = $('#tables li:first')
      $(table).addClass 'selected'

      showTableInfo()
      showTableStructure()

    cb()

showDatabaseInfo = () ->
  getInfo (data) ->
    $('#db_file_name').text(data.filename)
    $('#db_size').text(bytesToSize(data.size))
    $('#db_count_tables').text(data.number_of_tables)
    $('#db_count_indexes').text(data.number_of_indexes)

showTableInfo = () ->
  name = $('#tables li.selected').text()

  if name.length == 0
    alert 'No table selected. Please, select a table.'
    return

  getTableInfo name, (data) ->
    $('#table_information').show()
    $('#table_count_rows').text(data.row_count)

showTableStructure = () ->
  name = $('#tables li.selected').text()

  if name.length == 0
    alert 'No table selected. Please, select a table.'
    return

  buildTableStructure name, () ->
    setActiveTab 'table_structure'
    $('#structure').show()
    $('#input').hide()
    $('#output').hide()

showTableContent = () ->
  name = $('#tables li.selected').text()

  if name.length == 0
    alert 'No table selected. Please, select a table.'
    return

  buildTableContent name, () ->
    setActiveTab 'table_content'
    $('#structure').hide()
    $('#input').hide()
    $('#output').addClass 'full'
    $('#output').show()

showTableQuery = (query) ->
  resetResultTable()
  setActiveTab 'table_query'
  $('#structure').hide()
  $('#output').removeClass 'full'
  $('#input').show()
  $('#output').show()


####
## SQL query and export functions
####

runQuery = (query) ->
  $('#run #export_csv #export_json').prop 'disabled', true
  # TODO: Show progress

  buildTableQueryResult query
  $('#run #export_csv #export_json').prop 'disabled', false

exportCSV = (query) ->
  query = window.encodeURI(query.replace(/\n/g, ' '))
  host = window.location.host
  url = 'http://' + host + '/api/query?format=csv&query=' + query
  win = window.open(url, '_blank')

exportJSON = (query) ->
  query = window.encodeURI(query.replace(/\n/g, ' '))
  host = window.location.host
  url = 'http://' + host + '/api/query?format=json&query=' + query
  win = window.open(url, '_blank')


####
## Result table functions
####

buildResultHeader = (name) ->
  result = '<th>' + name + '</th>'

addHeadersToResultTable = (headers) ->
  header = '<thead><tr>'
  headers.forEach (h) ->
    header += h
  header += '</tr></thead>'

  $('#table_results').append(header)

buildResultRow = (row) ->
  result = '<tr>'
  row.forEach (v) ->
    result += '<th>' + v + '</th>'

  result += '</tr>'

addRowsToResultTable = (rows) ->
  body = '<tbody>'
  rows.forEach (row) ->
    body += row

  body += '</tbody>'
  $('#table_results').append(body)

resetResultTable = () ->
  $('#table_results').empty()


####
## UI helper functions
####

setActiveTab = (name) ->
  $('#navbar li.selected').removeClass 'selected'
  $('#' + name).addClass 'selected'


####
## Util functions
####

bytesToSize = (bytes) ->
  sizes = ['Bytes', 'KB', 'MB', 'GB', 'TB']
  if bytes == 0
    return '0 Byte'

  i = parseInt(Math.floor(Math.log(bytes) / Math.log(1024)))
  Math.round(bytes / Math.pow(1024, i), 2) + ' ' + sizes[i]

getColumnsFromIndexSql = (name, sql) ->
  i = sql.indexOf('"' + name + '"')
  sql = sql.slice(i+1, sql.length-1)
  matches = sql.match(/(\"[\w\s]+\")/g)
  result = (match.replace(/"/g, '') for match in matches)
  return result


####
## Main
####

$ ->
  # Setup Ace
  editor = ace.edit('editor')
  editor.setTheme('ace/theme/textmate')
  editor.getSession().setMode('ace/mode/sql')

  $('#tables').on 'click', 'li', () ->
    $('#tables li.selected').removeClass 'selected'
    $(this).addClass 'selected'

    showTableInfo()
    showTableStructure()

  $('#table_structure').on 'click', () ->
    showTableStructure()

  $('#table_content').on 'click', () ->
    showTableContent()

  $('#table_query').on 'click', () ->
    showTableQuery()

  # Setup index sql modal
  $('#index_sql_modal').on 'show.bs.modal', (event) ->
    button = $(event.relatedTarget)
    title = button.data('name')
    sql = button.next('pre').text()

    modal = $(this)
    modal.find('.modal-title').text(title)
    modal.find('.modal-body pre').text(sql)

  $('#run').on 'click', () ->
    query = $.trim editor.getValue()
    if query.length == 0
      return

    runQuery query

  $('#export_csv').on 'click', () ->
    query = $.trim editor.getValue()
    if query.length == 0
      return

    exportCSV query

  $('#export_json').on 'click', () ->
    query = $.trim editor.getValue()
    if query.length == 0
      return

    exportJSON query

  # Start here and load tables
  loadTables () ->
    showDatabaseInfo()
    $('#main').show()
