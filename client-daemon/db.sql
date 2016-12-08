create table ResampleRequest (
  id int auto_increment
  src VARCHAR(1024),
  dst VARCHAR(1024),
  dst_width int,
  dst_height int,
  dst_quality double,
  created_at datetime,
  updated_at datetime,
  status int, -- 0=enqueued, 1=in progress,2=done, 3=error
  message text -- error message
)
