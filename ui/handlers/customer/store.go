package customer

import (
	"fmt"

	"mpay/customer"
	"mpay/postgres"
	"mpay/postgres/querybuilder"
	"mpay/ui/components/table"
	"mpay/ui/pages/customer/consumer"
)

func selectNonDraftConsumersCount(db *postgres.Connection) (uint, error) {
	query := buildConsumersQuery()
	query.Columns("COUNT(*)")

	sql, args, err := query.String()
	if err != nil {
		fmt.Println("compose:", err)
		return 0, err
	}

	var count uint
	if err := db.QueryRow(sql, args...).Scan(&count); err != nil {
		fmt.Println("query: ", err)
		return 0, err
	}

	return count, nil
}

func selectAllNonDraftConsumers(db *postgres.Connection, offset, limit int) ([]consumer.ConsumerView, error) {
	query := buildConsumersQuery()
	query.Columns(
		"cus.id", "cus.code", "title.name_en", "cus.full_name", "cus.mobile_no",
		"cus.address_division_id", "division.name_en",
		"cus.address_township_id", "township.name_en",
		"cus.created_timestamp created_timestamp",
		"wallet.last_active_timestamp wallet_active_timestamp",
		"wallet.status wallet_status", "cus.status customer_status",
	)

	query.Limit(limit)
	query.Offset(offset)
	query.OrderBy("cus.created_timestamp DESC")

	sql, args, err := query.String()
	if err != nil {
		return nil, err
	}

	rows, err := db.Query(sql, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	results := make([]consumer.ConsumerView, 0, table.PAGE_SIZE)
	for rows.Next() {
		var v consumer.ConsumerView

		err := rows.Scan(
			&v.ID, &v.Code, &v.Title.Value, &v.FullName, &v.MobileNo,
			&v.DivisionID.Value, &v.DivisionName.Value,
			&v.TownshipID.Value, &v.TownshipName.Value,
			&v.CreatedTimestamp, &v.WalletActiveTimestamp.Value,
			&v.WalletStatus, &v.Status,
		)

		if err != nil {
			return nil, err
		}

		results = append(results, v)
	}

	return results, nil
}

func buildConsumersQuery() *querybuilder.SelectBuilder {
	query := querybuilder.NewSelectBuilder("customers AS cus")

	query.InnerJoin("customer_wallets AS wallet", "cus.id = wallet.customer_id")
	query.LeftJoin("reference_titles AS title", "cus.title = title.id")
	query.LeftJoin("reference_divisions AS division", "cus.address_division_id = division.id")
	query.LeftJoin("reference_townships AS township", "cus.address_township_id = township.id")

	where := query.Where("cus.type = ?", customer.Type_Consumer)
	where.And("cus.status != ?", customer.Status_Draft)

	return query
}

func selectNonDraftMerchantsCount(db *postgres.Connection) (uint, error) {
	query := buildMerchantsQuery()
	query.Columns("COUNT(*)")

	sql, args, err := query.String()
	if err != nil {
		return 0, err
	}

	var count uint
	if err := db.QueryRow(sql, args...).Scan(&count); err != nil {
		return 0, err
	}

	return count, nil
}

func selectAllNonDraftMerchants(db *postgres.Connection, offset, limit int) ([]consumer.MerchantView, error) {
	query := buildMerchantsQuery()
	query.Columns(
		"cus.id", "cus.code", "title.name_en", "cus.full_name",
		"parent.id parent_id", "parent.code parent_code",
		"parent_title.name_en parent_title", "parent.full_name parent_full_name",
		"cus.mobile_no",
		"cus.address_division_id", "division.name_en",
		"cus.address_township_id", "township.name_en",
		"cus.created_timestamp created_timestamp",
		"wallet.last_active_timestamp wallet_active_timestamp",
		"wallet.status wallet_status", "cus.status customer_status",
	)

	query.Limit(limit)
	query.Offset(offset)
	query.OrderBy("cus.created_timestamp DESC")

	sql, args, err := query.String()
	if err != nil {
		return nil, err
	}

	rows, err := db.Query(sql, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	results := make([]consumer.MerchantView, 0, table.PAGE_SIZE)
	for rows.Next() {
		var v consumer.MerchantView

		err := rows.Scan(
			&v.ID, &v.Code, &v.Title.Value, &v.FullName,
			&v.ParentID.Value, &v.ParentCode.Value,
			&v.ParentTitle.Value, &v.ParentFullName.Value,
			&v.MobileNo,
			&v.DivisionID.Value, &v.DivisionName.Value,
			&v.TownshipID.Value, &v.TownshipName.Value,
			&v.CreatedTimestamp, &v.WalletActiveTimestamp.Value,
			&v.WalletStatus, &v.Status,
		)

		if err != nil {
			return nil, err
		}

		results = append(results, v)
	}

	return results, nil
}

func buildMerchantsQuery() *querybuilder.SelectBuilder {
	query := querybuilder.NewSelectBuilder("customers AS cus")

	query.InnerJoin("customer_wallets AS wallet", "cus.id = wallet.customer_id")
	query.LeftJoin("customers AS parent", "cus.parent_id = parent.id")
	query.LeftJoin("reference_titles AS title", "cus.title = title.id")
	query.LeftJoin("reference_titles AS parent_title", "parent.title = parent_title.id")
	query.LeftJoin("reference_divisions AS division", "cus.address_division_id = division.id")
	query.LeftJoin("reference_townships AS township", "cus.address_township_id = township.id")

	merchantTypes := []any{
		customer.Type_MasterMerchant,
		customer.Type_IndividualMerchant,
		customer.Type_SubMerchant,
		customer.Type_Cashier,
	}
	where := query.Where("cus.type IN (?, ?, ?, ?)", merchantTypes...)
	where.And("cus.status != ?", customer.Status_Draft)

	return query
}

func selectNonDraftAgentsCount(db *postgres.Connection) (uint, error) {
	query := buildAgentsQuery()
	query.Columns("COUNT(*)")

	sql, args, err := query.String()
	if err != nil {
		return 0, err
	}

	var count uint
	if err := db.QueryRow(sql, args...).Scan(&count); err != nil {
		return 0, err
	}

	return count, nil
}

func selectAllNonDraftAgents(db *postgres.Connection, offset, limit int) ([]consumer.AgentView, error) {
	query := buildAgentsQuery()
	query.Columns(
		"cus.id", "cus.code", "title.name_en", "cus.full_name",
		"parent.id parent_id", "parent.code parent_code",
		"parent_title.name_en parent_title", "parent.full_name parent_full_name",
		"cus.mobile_no",
		"cus.address_division_id", "division.name_en",
		"cus.address_township_id", "township.name_en",
		"cus.created_timestamp created_timestamp",
		"wallet.last_active_timestamp wallet_active_timestamp",
		"wallet.status wallet_status", "cus.status customer_status",
	)

	query.Limit(limit)
	query.Offset(offset)
	query.OrderBy("cus.created_timestamp DESC")

	sql, args, err := query.String()
	if err != nil {
		return nil, err
	}

	rows, err := db.Query(sql, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	results := make([]consumer.AgentView, 0, table.PAGE_SIZE)
	for rows.Next() {
		var v consumer.AgentView

		err := rows.Scan(
			&v.ID, &v.Code, &v.Title.Value, &v.FullName,
			&v.ParentID.Value, &v.ParentCode.Value,
			&v.ParentTitle.Value, &v.ParentFullName.Value,
			&v.MobileNo,
			&v.DivisionID.Value, &v.DivisionName.Value,
			&v.TownshipID.Value, &v.TownshipName.Value,
			&v.CreatedTimestamp, &v.WalletActiveTimestamp.Value,
			&v.WalletStatus, &v.Status,
		)

		if err != nil {
			return nil, err
		}

		results = append(results, v)
	}

	return results, nil
}

func buildAgentsQuery() *querybuilder.SelectBuilder {
	query := querybuilder.NewSelectBuilder("customers AS cus")

	query.InnerJoin("customer_wallets AS wallet", "cus.id = wallet.customer_id")
	query.LeftJoin("customers AS parent", "cus.parent_id = parent.id")
	query.LeftJoin("reference_titles AS title", "cus.title = title.id")
	query.LeftJoin("reference_titles AS parent_title", "parent.title = parent_title.id")
	query.LeftJoin("reference_divisions AS division", "cus.address_division_id = division.id")
	query.LeftJoin("reference_townships AS township", "cus.address_township_id = township.id")

	agentTypes := []any{
		customer.Type_SuperAgent,
		customer.Type_VirtualAgent,
		customer.Type_Agent,
		customer.Type_SubAgent,
		customer.Type_CSE,
	}
	where := query.Where("cus.type IN (?, ?, ?, ?, ?)", agentTypes...)
	where.And("cus.status != ?", customer.Status_Draft)

	return query
}
