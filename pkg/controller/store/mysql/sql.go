package mysql

import "fmt"

const createUser = `
  insert into users (
    id,
    email,
    password_hash,
    first_name,
    last_name
  )
  values (?, ?, ?, ?, ?)
`

const getUser = `
  select id, email, password_hash, first_name, last_name, registration_completed from users
  where id = ?
`

const validateUser = `
  select id, email, password_hash, first_name, last_name, registration_completed from users
  where email = ? and password_hash = ?
`

const markRegistrationComplete = `
  update users
  set registration_completed = true
  where id = ?
`

const createRegistrationToken = `
  insert into registration_tokens (
    id,
    user_id,
    hash
  )
  values (?, ?, ?)
`

const getRegistrationToken = `
  select id, user_id, hash from registration_tokens
  where id = ?
`

const validateRegistrationToken = `
  select id, user_id, hash from registration_tokens
  where hash = ?
`

const createSession = `
  insert into sessions (
    id,
    user_id,
    hash
  )
  values (?, ?, ?)
`

const getSession = `
  select id, user_id, hash from sessions
  where id = ?
`

const validateSession = `
  select id, user_id, hash from sessions
  where hash = ?
`

const deleteSession = `
  delete from sessions
  where id = ?
  limit 1
`

const createAccessKey = `
  insert into access_keys (
    id,
    user_id,
    hash
  )
  values (?, ?, ?)
`

const getAccessKey = `
  select id, user_id, hash from access_keys
  where id = ?
`

const validateAccessKey = `
  select id, user_id, hash from access_keys
  where hash = ?
`

const createProject = `
  insert into projects (
    id,
    name
  )
  values (?, ?)
`

const getProject = `
  select id, name from projects
  where id = ?
`

const lookupProject = `
  select id, name from projects
  where name = ?
`

const getProjectDeviceCounts = `
  select count(*) from devices
  where project_id = ?
`

const getProjectApplicationCounts = `
  select count(*) from applications
  where project_id = ?
`

const createRole = `
  insert into roles (
    id,
    project_id,
    name,
    description,
    config
  )
  values (?, ?, ?, ?, ?)
`

const getRole = `
  select id, project_id, name, description, config from roles
  where id = ? and project_id = ?
`

const lookupRole = `
  select id, project_id, name, description, config from roles
  where name = ? and project_id = ?
`

const listRoles = `
  select id, project_id, name, description, config from roles
  where project_id = ?
`

const updateRole = `
  update roles
  set name = ?, description = ?, config = ?
  where id = ? and project_id = ?
`

const createMembership = `
  insert into memberships (
    id,
    user_id,
    project_id
  )
  values (?, ?, ?, ?)
`

const getMembership = `
  select id, user_id, project_id from memberships
  where user_id = ? and project_id = ?
`

const listMembershipsByUser = `
  select id, user_id, project_id from memberships
  where user_id = ?
`

const listMembershipsByProject = `
  select id, user_id, project_id from memberships
  where project_id = ?
`

const createMembershipRoleBinding = `
  insert into membership_role_bindings (
    membership_id,
    role_id,
    project_id
  )
  values (?, ?, ?)
`

const getMembershipRoleBinding = `
  select membership_id, role_id, project_id from membership_role_bindings
  where membership_id = ? and role_id = ? and project_id = ?
`

const listMembershipRoleBindings = `
  select membership_id, role_id, project_id from membership_role_bindings
  where membership_id = ? and project_id = ?
`

const createDevice = `
  insert into devices (
    id,
    project_id,
    name
  )
  values (?, ?, ?)
`

const getDevice = `
  select id, project_id, name, info from devices
  where id = ? and project_id = ?
`

const listDevices = `
  select id, project_id, name, info from devices
  where project_id = ?
`

const setDeviceInfo = `
  update devices
  set info = ?
  where id = ? and project_id = ?
`

var setDeviceLabel = fmt.Sprintf(`
  insert into device_labels (
    %skey%s,
    device_id,
    project_id,
    value
  )
  values (?, ?, ?, ?)
  on duplicate key update value = ?
`, "`", "`")

var getDeviceLabel = fmt.Sprintf(`
  select %skey%s, device_id, project_id, value from device_labels
  where %skey%s = ? and device_id = ? and project_id = ?
`, "`", "`", "`", "`")

var listDeviceLabels = fmt.Sprintf(`
  select %skey%s, device_id, project_id, value from device_labels
  where device_id = ? and project_id = ?
`, "`", "`")

var deleteDeviceLabel = fmt.Sprintf(`
  delete from device_labels
  where %skey%s = ? and device_id = ? and project_id = ?
  limit 1
`, "`", "`")

const createDeviceRegistrationToken = `
  insert into device_registration_tokens (
    id,
    project_id
  )
  values (?, ?)
`

const getDeviceRegistrationToken = `
  select id, project_id, device_access_key_id from device_registration_tokens
  where id = ? and project_id = ?
`

const bindDeviceRegistrationToken = `
  update device_registration_tokens
  set device_access_key_id = ?
  where id = ? and project_id = ?
`

const createDeviceAccessKey = `
  insert into device_access_keys (
    id,
    project_id,
    device_id,
    hash
  )
  values (?, ?, ?, ?)
`

const getDeviceAccessKey = `
  select id, project_id, device_id, hash from device_access_keys
  where id = ? and project_id = ?
`

const validateDeviceAccessKey = `
  select id, project_id, device_id, hash from device_access_keys
  where project_id = ? and hash = ?
`

const createApplication = `
  insert into applications (
    id,
    project_id,
    name
  )
  values (?, ?, ?)
`

const getApplication = `
  select id, project_id, name, settings from applications
  where id = ? and project_id = ?
`

const lookupApplication = `
  select id, project_id, name, settings from applications
  where name = ? and project_id = ?
`

const listApplications = `
  select id, project_id, name, settings from applications
  where project_id = ?
`

const setApplicationSettings = `
  update applications
  set settings = ?
  where id = ? and project_id = ?
`

const createRelease = `
  insert into releases (
    id,
    project_id,
    application_id,
    config
  )
  values (?, ?, ?, ?)
`

const getRelease = `
  select id, project_id, application_id, config from releases
  where id = ? and project_id = ? and application_id = ?
`

const getLatestRelease = `
  select id, project_id, application_id, config from releases
  where project_id = ? and application_id = ?
  order by created_at desc
  limit 1
`

const listReleases = `
  select id, project_id, application_id, config from releases
  where project_id = ? and application_id = ?
`

const getReleaseDeviceCounts = `
  select count(*) from device_application_statuses
  where project_id = ? and application_id = ? and current_release_id = ?
`

const setDeviceApplicationStatus = `
  insert into device_application_statuses (
    project_id,
    device_id,
    application_id,
    current_release_id
  )
  values (?, ?, ?, ?)
  on duplicate key update
    current_release_id = ?
`

const getDeviceApplicationStatus = `
  select project_id, device_id, application_id, current_release_id from device_application_statuses
  where project_id = ? and device_id = ? and application_id = ?
`

const setDeviceServiceStatus = `
  insert into device_service_statuses (
    project_id,
    device_id,
    application_id,
    service,
    current_release_id
  )
  values (?, ?, ?, ?, ?)
  on duplicate key update
    current_release_id = ?
`

const getDeviceServiceStatus = `
  select project_id, device_id, application_id, service, current_release_id from device_service_statuses
  where project_id = ? and device_id = ? and application_id = ? and service = ?
`

const getDeviceServiceStatuses = `
  select project_id, device_id, application_id, service, current_release_id from device_service_statuses
  where project_id = ? and device_id = ? and application_id = ?
`
