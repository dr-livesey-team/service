--[[---------------------------------------------------------------------------
Selectors
--]]---------------------------------------------------------------------------

function table_contains(table, value)
    if table == nil then
        return true
    end

    local contains = false
    for i, v in ipairs(table) do
        if v == value then
            contains = true
        end
    end
    return contains
end

function select_anomalies(filters)
    return box.space.anomalies
        :pairs()
        :filter(
            function (tuple)
                if tuple.opening_date < filters.opening_date then
                    return false
                end

                if tuple.closing_date > filters.closing_date then
                    return false
                end

                if not table_contains(filters.district_name, tuple.district_name) then
                    return false
                end

                if not table_contains(filters.address, tuple.address) then
                    return false
                end

                if not table_contains(filters.management_company_name, tuple.address) then
                    return false
                end

                if not table_contains(filters.service_organization_name, tuple.service_organization_name) then
                    return false
                end

                if not table_contains(filters.urgency_category_name, tuple.urgency_category_name) then
                    return false
                end

                if not table_contains(filters.anomaly_category, tuple.anomaly_category) then
                    return false
                end

                return true
            end
        )
        -- :map(
        --     function (tuple)
        --         return {tuple.id, tuple.latitude, tuple.longitude}
        --     end
        -- )
        :totable()
end

function number_anomalies(filters)
    return box.space.anomalies
        :pairs()
        :filter(
            function(tuple)
                if tuple.opening_date < filters.opening_date then
                    return false
                end

                if tuple.closing_date > filters.closing_date then
                    return false
                end

                if not table_contains(filters.district_name, tuple.district_name) then
                    return false
                end

                if not table_contains(filters.address, tuple.address) then
                    return false
                end

                if not table_contains(filters.management_company_name, tuple.address) then
                    return false
                end

                if not table_contains(filters.service_organization_name, tuple.service_organization_name) then
                    return false
                end

                if not table_contains(filters.urgency_category_name, tuple.urgency_category_name) then
                    return false
                end

                if not table_contains(filters.anomaly_category, tuple.anomaly_category) then
                    return false
                end

                return true
            end
        )
        :length()
end

function number_normal(filters)
    return box.space.normal
        :pairs()
        :filter(
            function (tuple)
                if tuple.opening_date < filters.opening_date then
                    return false
                end

                if tuple.closing_date > filters.closing_date then
                    return false
                end

                if not table_contains(filters.district_name, tuple.district_name) then
                    return false
                end

                if not table_contains(filters.address, tuple.address) then
                    return false
                end

                if not table_contains(filters.management_company_name, tuple.address) then
                    return false
                end

                if not table_contains(filters.service_organization_name, tuple.service_organization_name) then
                    return false
                end

                if not table_contains(filters.urgency_category_name, tuple.urgency_category_name) then
                    return false
                end

                return true
            end
        )
        :length()
end

function select_requests(id)
    return box.space.requests
        :pairs()
        :filter(
            function (tuple)
                return tuple.group_id == id
            end
        )
        :totable()
end

--[[---------------------------------------------------------------------------
Configure tarantool
--]]---------------------------------------------------------------------------

box.cfg{
    background = true,
    listen = '127.0.0.1:3301'
}

--[[---------------------------------------------------------------------------
Load necessary lua modules
--]]---------------------------------------------------------------------------

fio = require('fio')

csv = require('csv')

--[[---------------------------------------------------------------------------
Set up tarantool remote user options
--]]---------------------------------------------------------------------------

box.schema.user.passwd('passwd')

--[[---------------------------------------------------------------------------
Creation of space for anomaly situations (situation is set of request grouped by
belonging to one anomaly)
--]]---------------------------------------------------------------------------

if box.space.anomalies == nil then
    box.schema.space.create('anomalies')

    box.space.anomalies:format({
        {name = 'id',                        type = 'unsigned'},
        {name = 'opening_date',              type = 'string'  },
        {name = 'closing_date',              type = 'string'  },
        {name = 'district_name',             type = 'string'  },
        {name = 'address',                   type = 'string'  },
        {name = 'fault_name',                type = 'string'  },
        {name = 'management_company_name',   type = 'string'  },
        {name = 'service_organization_name', type = 'string'  },
        {name = 'urgency_category_name',     type = 'string'  },
        {name = 'anomaly_category',          type = 'string'  },
        {name = 'latitude',                  type = 'number'  },
        {name = 'longitude',                 type = 'number'  }
    })

    box.space.anomalies:create_index('primary', {parts = {'id'}})
end

--[[---------------------------------------------------------------------------
Creation of anomalies selector
--]]---------------------------------------------------------------------------

if box.func.select_anomalies == nil then
    box.schema.func.create('select_anomalies')
end
if box.func.number_anomalies == nil then
    box.schema.func.create('number_anomalies')
end
if box.func.select_requests == nil then
    box.schema.func.create('select_requests')
end

--[[---------------------------------------------------------------------------
Creation of space for normal situations (situation is set of request grouped by
belonging to one anomaly)
--]]---------------------------------------------------------------------------

if box.space.normal == nil then
    box.schema.space.create('normal')

    box.space.normal:format({
        {name = 'id',                        type = 'unsigned'},
        {name = 'opening_date',              type = 'string'  },
        {name = 'closing_date',              type = 'string'  },
        {name = 'district_name',             type = 'string'  },
        {name = 'address',                   type = 'string'  },
        {name = 'fault_name',                type = 'string'  },
        {name = 'management_company_name',   type = 'string'  },
        {name = 'service_organization_name', type = 'string'  },
        {name = 'urgency_category_name',     type = 'string'  }
    })

    box.space.normal:create_index('primary', {parts = {'id'}})
end

--[[---------------------------------------------------------------------------
Creation of anomalies selector
--]]---------------------------------------------------------------------------

if box.func.number_normal == nil then
    box.schema.func.create('number_normal')
end

--[[---------------------------------------------------------------------------
Creation of space for all request data
--]]---------------------------------------------------------------------------

if box.space.requests == nil then
    box.schema.space.create('requests')

    box.space.requests:format({
        {name = 'request_root_identifier',   type = 'unsigned'},
        {name = 'opening_date',              type = 'string'  },
        {name = 'closing_date',              type = 'string'  },
        {name = 'district_name',             type = 'string'  },
        {name = 'address',                   type = 'string'  },
        {name = 'fault_name',                type = 'string'  },
        {name = 'management_company_name',   type = 'string'  },
        {name = 'service_organization_name', type = 'string'  },
        {name = 'urgency_category_name',     type = 'string'  },
        {name = 'feedback',                  type = 'string'  },
        {name = 'group_id',                  type = 'unsigned'}
    })

    box.space.requests:create_index('primary', {parts = {'request_root_identifier'}})
end

--[[---------------------------------------------------------------------------
Creation of requests data selector
--]]---------------------------------------------------------------------------

if box.func.select_requests == nil then
    box.schema.func.create('select_requests')
end