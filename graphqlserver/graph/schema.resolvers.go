package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.42

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/intelops/kubviz/graphqlserver/graph/model"
)

// AllEvents is the resolver for the allEvents field.
func (r *queryResolver) AllEvents(ctx context.Context) ([]*model.Event, error) {
	if r.DB == nil {
		return nil, fmt.Errorf("database connection is not initialized")
	}

	query := `SELECT ClusterName, Id, EventTime, OpType, Name, Namespace, Kind, Message, Reason, Host, Event, FirstTime, LastTime, ExpiryDate FROM events`

	rows, err := r.DB.QueryContext(ctx, query)
	if err != nil {
		if err == sql.ErrNoRows {
			return []*model.Event{}, nil
		}
		return nil, fmt.Errorf("error executing query: %v", err)
	}
	defer rows.Close()

	var events []*model.Event
	for rows.Next() {
		var e model.Event
		if err := rows.Scan(&e.ClusterName, &e.ID, &e.EventTime, &e.OpType, &e.Name, &e.Namespace, &e.Kind, &e.Message, &e.Reason, &e.Host, &e.Event, &e.FirstTime, &e.LastTime, &e.ExpiryDate); err != nil {
			return nil, fmt.Errorf("error scanning row: %v", err)
		}
		events = append(events, &e)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %v", err)
	}

	return events, nil
}

// AllRakkess is the resolver for the allRakkess field.
func (r *queryResolver) AllRakkess(ctx context.Context) ([]*model.Rakkess, error) {
	if r.DB == nil {
		return nil, fmt.Errorf("database connection is not initialized")
	}

	query := `SELECT ClusterName, Name, Create, Delete, List, Update, EventTime, ExpiryDate FROM rakkess`

	rows, err := r.DB.QueryContext(ctx, query)
	if err != nil {
		if err == sql.ErrNoRows {
			return []*model.Rakkess{}, nil
		}
		return nil, fmt.Errorf("error executing query: %v", err)
	}
	defer rows.Close()

	var rakkessRecords []*model.Rakkess
	for rows.Next() {
		var r model.Rakkess
		if err := rows.Scan(&r.ClusterName, &r.Name, &r.Create, &r.Delete, &r.List, &r.Update, &r.EventTime, &r.ExpiryDate); err != nil {
			return nil, fmt.Errorf("error scanning row: %v", err)
		}
		rakkessRecords = append(rakkessRecords, &r)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %v", err)
	}

	return rakkessRecords, nil
}

// AllDeprecatedAPIs is the resolver for the allDeprecatedAPIs field.
func (r *queryResolver) AllDeprecatedAPIs(ctx context.Context) ([]*model.DeprecatedAPI, error) {
	if r.DB == nil {
		return nil, fmt.Errorf("database connection is not initialized")
	}

	query := `SELECT ClusterName, ObjectName, Description, Kind, Deprecated, Scope, EventTime, ExpiryDate FROM DeprecatedAPIs`

	rows, err := r.DB.QueryContext(ctx, query)
	if err != nil {
		if err == sql.ErrNoRows {
			return []*model.DeprecatedAPI{}, nil
		}
		return nil, fmt.Errorf("error executing query: %v", err)
	}
	defer rows.Close()

	var deprecatedAPIs []*model.DeprecatedAPI
	for rows.Next() {
		var d model.DeprecatedAPI
		var deprecatedInt uint8
		if err := rows.Scan(&d.ClusterName, &d.ObjectName, &d.Description, &d.Kind, &deprecatedInt, &d.Scope, &d.EventTime, &d.ExpiryDate); err != nil {
			return nil, fmt.Errorf("error scanning row: %v", err)
		}

		// Convert uint8 to bool
		deprecatedBool := deprecatedInt != 0
		d.Deprecated = &deprecatedBool

		deprecatedAPIs = append(deprecatedAPIs, &d)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %v", err)
	}

	return deprecatedAPIs, nil
}

// AllDeletedAPIs is the resolver for the allDeletedAPIs field.
func (r *queryResolver) AllDeletedAPIs(ctx context.Context) ([]*model.DeletedAPI, error) {
	if r.DB == nil {
		return nil, fmt.Errorf("database connection is not initialized")
	}

	query := `SELECT ClusterName, ObjectName, Group, Kind, Version, Name, Deleted, Scope, EventTime, ExpiryDate FROM DeletedAPIs`

	rows, err := r.DB.QueryContext(ctx, query)
	if err != nil {
		if err == sql.ErrNoRows {
			return []*model.DeletedAPI{}, nil
		}
		return nil, fmt.Errorf("error executing query: %v", err)
	}
	defer rows.Close()

	var deletedAPIs []*model.DeletedAPI
	for rows.Next() {
		var d model.DeletedAPI
		var deletedInt uint8
		if err := rows.Scan(&d.ClusterName, &d.ObjectName, &d.Group, &d.Kind, &d.Version, &d.Name, &deletedInt, &d.Scope, &d.EventTime, &d.ExpiryDate); err != nil {
			return nil, fmt.Errorf("error scanning row: %v", err)
		}

		// Convert uint8 to bool
		deletedBool := deletedInt != 0
		d.Deleted = &deletedBool

		deletedAPIs = append(deletedAPIs, &d)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %v", err)
	}

	return deletedAPIs, nil
}

// AllGetAllResources is the resolver for the allGetAllResources field.
func (r *queryResolver) AllGetAllResources(ctx context.Context) ([]*model.GetAllResource, error) {
	if r.DB == nil {
		return nil, fmt.Errorf("database connection is not initialized")
	}

	query := `SELECT ClusterName, Namespace, Kind, Resource, Age, EventTime, ExpiryDate FROM getall_resources`

	rows, err := r.DB.QueryContext(ctx, query)
	if err != nil {
		if err == sql.ErrNoRows {
			return []*model.GetAllResource{}, nil
		}
		return nil, fmt.Errorf("error executing query: %v", err)
	}
	defer rows.Close()

	var resources []*model.GetAllResource
	for rows.Next() {
		var res model.GetAllResource
		if err := rows.Scan(&res.ClusterName, &res.Namespace, &res.Kind, &res.Resource, &res.Age, &res.EventTime, &res.ExpiryDate); err != nil {
			return nil, fmt.Errorf("error scanning row: %v", err)
		}
		resources = append(resources, &res)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %v", err)
	}

	return resources, nil
}

// AllTrivySBOMs is the resolver for the allTrivySBOMs field.
func (r *queryResolver) AllTrivySBOMs(ctx context.Context) ([]*model.TrivySbom, error) {
	if r.DB == nil {
		return nil, fmt.Errorf("database connection is not initialized")
	}

	query := `SELECT id, cluster_name, image_name, package_name, package_url, bom_ref, serial_number, version, bom_format, ExpiryDate FROM trivysbom`

	rows, err := r.DB.QueryContext(ctx, query)
	if err != nil {
		if err == sql.ErrNoRows {
			return []*model.TrivySbom{}, nil
		}
		return nil, fmt.Errorf("error executing query: %v", err)
	}
	defer rows.Close()

	var sboms []*model.TrivySbom
	for rows.Next() {
		var s model.TrivySbom
		if err := rows.Scan(&s.ID, &s.ClusterName, &s.ImageName, &s.PackageName, &s.PackageURL, &s.BomRef, &s.SerialNumber, &s.Version, &s.BomFormat, &s.ExpiryDate); err != nil {
			return nil, fmt.Errorf("error scanning row: %v", err)
		}
		sboms = append(sboms, &s)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %v", err)
	}

	return sboms, nil
}

// AllTrivyImages is the resolver for the allTrivyImages field.
func (r *queryResolver) AllTrivyImages(ctx context.Context) ([]*model.TrivyImage, error) {
	if r.DB == nil {
		return nil, fmt.Errorf("database connection is not initialized")
	}

	query := `SELECT id, cluster_name, artifact_name, vul_id, vul_pkg_id, vul_pkg_name, vul_installed_version, vul_fixed_version, vul_title, vul_severity, vul_published_date, vul_last_modified_date, ExpiryDate FROM trivyimage`

	rows, err := r.DB.QueryContext(ctx, query)
	if err != nil {
		if err == sql.ErrNoRows {
			return []*model.TrivyImage{}, nil
		}
		return nil, fmt.Errorf("error executing query: %v", err)
	}
	defer rows.Close()

	var images []*model.TrivyImage
	for rows.Next() {
		var img model.TrivyImage
		if err := rows.Scan(&img.ID, &img.ClusterName, &img.ArtifactName, &img.VulID, &img.VulPkgID, &img.VulPkgName, &img.VulInstalledVersion, &img.VulFixedVersion, &img.VulTitle, &img.VulSeverity, &img.VulPublishedDate, &img.VulLastModifiedDate, &img.ExpiryDate); err != nil {
			return nil, fmt.Errorf("error scanning row: %v", err)
		}
		images = append(images, &img)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %v", err)
	}

	return images, nil
}

// AllKubeScores is the resolver for the allKubeScores field.
func (r *queryResolver) AllKubeScores(ctx context.Context) ([]*model.Kubescore, error) {
	if r.DB == nil {
		return nil, fmt.Errorf("database connection is not initialized")
	}

	query := `SELECT id, clustername, object_name, kind, apiVersion, name, namespace, target_type, description, path, summary, file_name, file_row, EventTime, ExpiryDate FROM kubescore`

	rows, err := r.DB.QueryContext(ctx, query)
	if err != nil {
		if err == sql.ErrNoRows {
			return []*model.Kubescore{}, nil
		}
		return nil, fmt.Errorf("error executing query: %v", err)
	}
	defer rows.Close()

	var kubeScores []*model.Kubescore
	for rows.Next() {
		var ks model.Kubescore
		if err := rows.Scan(&ks.ID, &ks.ClusterName, &ks.ObjectName, &ks.Kind, &ks.APIVersion, &ks.Name, &ks.Namespace, &ks.TargetType, &ks.Description, &ks.Path, &ks.Summary, &ks.FileName, &ks.FileRow, &ks.EventTime); err != nil {
			return nil, fmt.Errorf("error scanning row: %v", err)
		}
		kubeScores = append(kubeScores, &ks)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %v", err)
	}

	return kubeScores, nil
}

// AllTrivyVuls is the resolver for the allTrivyVuls field.
func (r *queryResolver) AllTrivyVuls(ctx context.Context) ([]*model.TrivyVul, error) {
	if r.DB == nil {
		return nil, fmt.Errorf("database connection is not initialized")
	}

	query := `SELECT id, cluster_name, namespace, kind, name, vul_id, vul_vendor_ids, vul_pkg_id, vul_pkg_name, vul_pkg_path, vul_installed_version, vul_fixed_version, vul_title, vul_severity, vul_published_date, vul_last_modified_date, ExpiryDate FROM trivy_vul`

	rows, err := r.DB.QueryContext(ctx, query)
	if err != nil {
		if err == sql.ErrNoRows {
			return []*model.TrivyVul{}, nil
		}
		return nil, fmt.Errorf("error executing query: %v", err)
	}
	defer rows.Close()

	var trivyVuls []*model.TrivyVul
	for rows.Next() {
		var tv model.TrivyVul
		if err := rows.Scan(&tv.ID, &tv.ClusterName, &tv.Namespace, &tv.Kind, &tv.Name, &tv.VulID, &tv.VulVendorIds, &tv.VulPkgID, &tv.VulPkgName, &tv.VulPkgPath, &tv.VulInstalledVersion, &tv.VulFixedVersion, &tv.VulTitle, &tv.VulSeverity, &tv.VulPublishedDate, &tv.VulLastModifiedDate, &tv.ExpiryDate); err != nil {
			return nil, fmt.Errorf("error scanning row: %v", err)
		}
		trivyVuls = append(trivyVuls, &tv)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %v", err)
	}

	return trivyVuls, nil
}

// AllTrivyMisconfigs is the resolver for the allTrivyMisconfigs field.
func (r *queryResolver) AllTrivyMisconfigs(ctx context.Context) ([]*model.TrivyMisconfig, error) {
	if r.DB == nil {
		return nil, fmt.Errorf("database connection is not initialized")
	}

	query := `SELECT id, cluster_name, namespace, kind, name, misconfig_id, misconfig_avdid, misconfig_type, misconfig_title, misconfig_desc, misconfig_msg, misconfig_query, misconfig_resolution, misconfig_severity, misconfig_status, EventTime, ExpiryDate FROM trivy_misconfig`

	rows, err := r.DB.QueryContext(ctx, query)
	if err != nil {
		if err == sql.ErrNoRows {
			return []*model.TrivyMisconfig{}, nil
		}
		return nil, fmt.Errorf("error executing query: %v", err)
	}
	defer rows.Close()

	var misconfigs []*model.TrivyMisconfig
	for rows.Next() {
		var tm model.TrivyMisconfig
		if err := rows.Scan(&tm.ID, &tm.ClusterName, &tm.Namespace, &tm.Kind, &tm.Name, &tm.MisconfigID, &tm.MisconfigAvdid, &tm.MisconfigType, &tm.MisconfigTitle, &tm.MisconfigDesc, &tm.MisconfigMsg, &tm.MisconfigQuery, &tm.MisconfigResolution, &tm.MisconfigSeverity, &tm.MisconfigStatus, &tm.EventTime, &tm.ExpiryDate); err != nil {
			return nil, fmt.Errorf("error scanning row: %v", err)
		}
		misconfigs = append(misconfigs, &tm)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %v", err)
	}

	return misconfigs, nil
}

// UniqueNamespaces is the resolver for the uniqueNamespaces field.
func (r *queryResolver) UniqueNamespaces(ctx context.Context, clusterName string) ([]*model.Namespace, error) {
	namespaces, err := r.fetchNamespacesFromDatabase(ctx, clusterName)
	if err != nil {
		return nil, err
	}

	var namespaceObjects []*model.Namespace
	for _, ns := range namespaces {
		namespaceObjects = append(namespaceObjects, &model.Namespace{Name: ns})
	}

	return namespaceObjects, nil
}

// UniqueClusters is the resolver for the uniqueClusters field.
func (r *queryResolver) UniqueClusters(ctx context.Context) ([]*model.Cluster, error) {
	clusters, err := r.fetchClustersFromDatabase(ctx)
	if err != nil {
		return nil, err
	}

	var clusterObjects []*model.Cluster
	for _, cluster := range clusters {
		clusterObjects = append(clusterObjects, &model.Cluster{Name: cluster})
	}

	return clusterObjects, nil
}

// OutdatedImagesByClusterAndNamespace is the resolver for the outdatedImagesByClusterAndNamespace field.
func (r *queryResolver) OutdatedImagesByClusterAndNamespace(ctx context.Context, clusterName string, namespace string) ([]*model.OutdatedImage, error) {
	if r.DB == nil {
		return nil, fmt.Errorf("database connection is not initialized")
	}

	if clusterName == "" || namespace == "" {
		return nil, fmt.Errorf("clusterName and namespace cannot be empty")
	}

	query := `SELECT ClusterName, Namespace, Pod, CurrentImage, CurrentTag, LatestVersion, VersionsBehind, EventTime FROM outdated_images WHERE ClusterName = ? AND Namespace = ?`

	rows, err := r.DB.QueryContext(ctx, query, clusterName, namespace)
	if err != nil {
		if err == sql.ErrNoRows {
			return []*model.OutdatedImage{}, nil
		}
		return nil, fmt.Errorf("error executing query: %v", err)
	}
	defer rows.Close()

	var outdatedImages []*model.OutdatedImage
	for rows.Next() {
		var oi model.OutdatedImage
		if err := rows.Scan(&oi.ClusterName, &oi.Namespace, &oi.Pod, &oi.CurrentImage, &oi.CurrentTag, &oi.LatestVersion, &oi.VersionsBehind, &oi.EventTime); err != nil {
			return nil, fmt.Errorf("error scanning row: %v", err)
		}
		outdatedImages = append(outdatedImages, &oi)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %v", err)
	}

	return outdatedImages, nil
}

// OutdatedImagesCount is the resolver for the outdatedImagesCount field.
func (r *queryResolver) OutdatedImagesCount(ctx context.Context, clusterName string, namespace string) (int, error) {
	if r.DB == nil {
		return 0, fmt.Errorf("database connection is not initialized")
	}

	if clusterName == "" || namespace == "" {
		return 0, fmt.Errorf("clusterName and namespace cannot be empty")
	}

	query := `SELECT COUNT(*) FROM outdated_images WHERE ClusterName = ? AND Namespace = ?`

	var count int
	err := r.DB.QueryRowContext(ctx, query, clusterName, namespace).Scan(&count)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, nil
		}
		return 0, fmt.Errorf("error executing query: %v", err)
	}

	return count, nil
}

// AllClusterNamespaceOutdatedCounts is the resolver for the allClusterNamespaceOutdatedCounts field.
func (r *queryResolver) AllClusterNamespaceOutdatedCounts(ctx context.Context) ([]*model.ClusterNamespaceOutdatedCount, error) {
	if r.DB == nil {
		return nil, fmt.Errorf("database connection is not initialized")
	}

	query := `
        SELECT ClusterName, Namespace, COUNT(*) as outdatedCount
        FROM outdated_images
        GROUP BY ClusterName, Namespace
    `

	rows, err := r.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %v", err)
	}
	defer rows.Close()

	var results []*model.ClusterNamespaceOutdatedCount
	for rows.Next() {
		var result model.ClusterNamespaceOutdatedCount
		if err := rows.Scan(&result.ClusterName, &result.Namespace, &result.OutdatedCount); err != nil {
			return nil, fmt.Errorf("error scanning row: %v", err)
		}
		results = append(results, &result)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %v", err)
	}

	return results, nil
}

// AllClusterDeprecatedAPIsCounts is the resolver for the allClusterDeprecatedAPIsCounts field.
func (r *queryResolver) AllClusterDeprecatedAPIsCounts(ctx context.Context) ([]*model.ClusterAPIsCount, error) {
	if r.DB == nil {
		return nil, fmt.Errorf("database connection is not initialized")
	}
	query := `
	SELECT ClusterName, COUNT(*) as count
	FROM DeprecatedAPIs
	GROUP BY ClusterName
`

	rows, err := r.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %v", err)
	}
	defer rows.Close()
	var results []*model.ClusterAPIsCount
	for rows.Next() {
		var result model.ClusterAPIsCount
		if err := rows.Scan(&result.ClusterName, &result.Count); err != nil {
			return nil, fmt.Errorf("error scanning row: %v", err)
		}
		results = append(results, &result)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %v", err)
	}

	return results, nil
}

// AllClusterDeletedAPIsCounts is the resolver for the allClusterDeletedAPIsCounts field.
func (r *queryResolver) AllClusterDeletedAPIsCounts(ctx context.Context) ([]*model.ClusterAPIsCount, error) {
	if r.DB == nil {
		return nil, fmt.Errorf("database connection is not initialized")
	}

	query := `
        SELECT ClusterName, COUNT(*) as count
        FROM DeletedAPIs
        GROUP BY ClusterName
    `

	rows, err := r.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %v", err)
	}
	defer rows.Close()

	var results []*model.ClusterAPIsCount
	for rows.Next() {
		var result model.ClusterAPIsCount
		if err := rows.Scan(&result.ClusterName, &result.Count); err != nil {
			return nil, fmt.Errorf("error scanning row: %v", err)
		}
		results = append(results, &result)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %v", err)
	}

	return results, nil
}

// AllClusterNamespaceResourceCounts is the resolver for the allClusterNamespaceResourceCounts field.
func (r *queryResolver) AllClusterNamespaceResourceCounts(ctx context.Context) ([]*model.ClusterNamespaceResourceCount, error) {
	if r.DB == nil {
		return nil, fmt.Errorf("database connection is not initialized")
	}

	query := `
        SELECT ClusterName, Namespace, COUNT(*) as resourceCount
        FROM getall_resources
        GROUP BY ClusterName, Namespace
    `

	rows, err := r.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %v", err)
	}
	defer rows.Close()

	var results []*model.ClusterNamespaceResourceCount
	for rows.Next() {
		var result model.ClusterNamespaceResourceCount
		if err := rows.Scan(&result.ClusterName, &result.Namespace, &result.ResourceCount); err != nil {
			return nil, fmt.Errorf("error scanning row: %v", err)
		}
		results = append(results, &result)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %v", err)
	}

	return results, nil
}

// EventsByClusterAndNamespace is the resolver for the eventsByClusterAndNamespace field.
func (r *queryResolver) EventsByClusterAndNamespace(ctx context.Context, clusterName string, namespace string) ([]*model.Event, error) {
	if r.DB == nil {
		return nil, fmt.Errorf("database connection is not initialized")
	}

	if clusterName == "" || namespace == "" {
		return nil, fmt.Errorf("clusterName and namespace cannot be empty")
	}

	query := `SELECT ClusterName, Id, EventTime, OpType, Name, Namespace, Kind, Message, Reason, Host, Event, ImageName, FirstTime, LastTime FROM events WHERE ClusterName = ? AND Namespace = ?`

	rows, err := r.DB.QueryContext(ctx, query, clusterName, namespace)
	if err != nil {
		if err == sql.ErrNoRows {
			return []*model.Event{}, nil
		}
		return nil, fmt.Errorf("error executing query: %v", err)
	}
	defer rows.Close()

	var events []*model.Event
	for rows.Next() {
		var e model.Event
		if err := rows.Scan(&e.ClusterName, &e.ID, &e.EventTime, &e.OpType, &e.Name, &e.Namespace, &e.Kind, &e.Message, &e.Reason, &e.Host, &e.Event, &e.ImageName, &e.FirstTime, &e.LastTime); err != nil {
			return nil, fmt.Errorf("error scanning row: %v", err)
		}
		events = append(events, &e)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %v", err)
	}

	return events, nil
}

// Vulnerabilities is the resolver for the vulnerabilities field.
func (r *queryResolver) Vulnerabilities(ctx context.Context, clusterName string, namespace string) ([]*model.Vulnerability, error) {
	if r.DB == nil {
		return nil, fmt.Errorf("database connection is not initialized")
	}

	if clusterName == "" || namespace == "" {
		return nil, fmt.Errorf("clusterName and namespace cannot be empty")
	}
	query := `
	SELECT id, cluster_name, namespace, kind, name, vul_id, vul_vendor_ids, vul_pkg_id, vul_pkg_name, vul_pkg_path, vul_installed_version, vul_fixed_version, vul_title, vul_severity, vul_published_date, vul_last_modified_date, ExpiryDate
	FROM trivy_vul
	WHERE cluster_name = ? AND namespace = ?
	`
	rows, err := r.DB.QueryContext(ctx, query, clusterName, namespace)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %v", err)
	}
	defer rows.Close()
	var vulnerabilities []*model.Vulnerability
	for rows.Next() {
		var v model.Vulnerability
		if err := rows.Scan(&v.ID, &v.ClusterName, &v.Namespace, &v.Kind, &v.Name, &v.VulID, &v.VulVendorIds, &v.VulPkgID, &v.VulPkgName, &v.VulPkgPath, &v.VulInstalledVersion, &v.VulFixedVersion, &v.VulTitle, &v.VulSeverity, &v.VulPublishedDate, &v.VulLastModifiedDate, &v.ExpiryDate, &v.ExportedAt); err != nil {
			return nil, fmt.Errorf("error scanning row: %v", err)
		}
		vulnerabilities = append(vulnerabilities, &v)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %v", err)
	}

	return vulnerabilities, nil
}

// Misconfigurations is the resolver for the misconfigurations field.
func (r *queryResolver) Misconfigurations(ctx context.Context, clusterName string, namespace string) ([]*model.Misconfiguration, error) {
	if r.DB == nil {
		return nil, fmt.Errorf("database connection is not initialized")
	}

	if clusterName == "" || namespace == "" {
		return nil, fmt.Errorf("clusterName and namespace cannot be empty")
	}

	query := `
	SELECT id, cluster_name, namespace, kind, name, misconfig_id, misconfig_avdid, misconfig_type, misconfig_title, misconfig_desc, misconfig_msg, misconfig_query, misconfig_resolution, misconfig_severity, misconfig_status, EventTime, ExpiryDate
	FROM trivy_misconfig
	WHERE cluster_name = ? AND namespace = ?
	`

	rows, err := r.DB.QueryContext(ctx, query, clusterName, namespace)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %v", err)
	}
	defer rows.Close()

	var misconfigurations []*model.Misconfiguration
	for rows.Next() {
		var m model.Misconfiguration
		if err := rows.Scan(&m.ID, &m.ClusterName, &m.Namespace, &m.Kind, &m.Name, &m.MisconfigID, &m.MisconfigAvdid, &m.MisconfigType, &m.MisconfigTitle, &m.MisconfigDesc, &m.MisconfigMsg, &m.MisconfigQuery, &m.MisconfigResolution, &m.MisconfigSeverity, &m.MisconfigStatus, &m.EventTime, &m.ExpiryDate, &m.ExportedAt); err != nil {
			return nil, fmt.Errorf("error scanning row: %v", err)
		}
		misconfigurations = append(misconfigurations, &m)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %v", err)
	}

	return misconfigurations, nil
}

// Kubescores is the resolver for the kubescores field.
func (r *queryResolver) Kubescores(ctx context.Context, clustername string, namespace string) ([]*model.KubeScore, error) {
	if r.DB == nil {
		return nil, fmt.Errorf("database connection is not initialized")
	}

	if clustername == "" || namespace == "" {
		return nil, fmt.Errorf("clustername and namespace cannot be empty")
	}

	query := `
	SELECT id, clustername, object_name, kind, apiVersion, name, namespace, target_type, description, path, summary, file_name, file_row, EventTime, ExpiryDate
	FROM kubescore
	WHERE clustername = ? AND namespace = ?
	`

	rows, err := r.DB.QueryContext(ctx, query, clustername, namespace)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %v", err)
	}
	defer rows.Close()

	var scores []*model.KubeScore
	for rows.Next() {
		var s model.KubeScore
		if err := rows.Scan(&s.ID, &s.ClusterName, &s.ObjectName, &s.Kind, &s.APIVersion, &s.Name, &s.Namespace, &s.TargetType, &s.Description, &s.Path, &s.Summary, &s.FileName, &s.FileRow, &s.EventTime); err != nil {
			return nil, fmt.Errorf("error scanning row: %v", err)
		}
		scores = append(scores, &s)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %v", err)
	}

	return scores, nil
}

// GetAllResources is the resolver for the getAllResources field.
func (r *queryResolver) GetAllResources(ctx context.Context, clusterName string, namespace string) ([]*model.GetAllResource, error) {
	if r.DB == nil {
		return nil, fmt.Errorf("database connection is not initialized")
	}

	if clusterName == "" || namespace == "" {
		return nil, fmt.Errorf("clusterName and namespace cannot be empty")
	}

	query := `
	SELECT ClusterName, Namespace, Kind, Resource, Age, EventTime, ExpiryDate
	FROM getall_resources
	WHERE ClusterName = ? AND Namespace = ?
	`

	rows, err := r.DB.QueryContext(ctx, query, clusterName, namespace)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %v", err)
	}
	defer rows.Close()

	var resources []*model.GetAllResource
	for rows.Next() {
		var r model.GetAllResource
		if err := rows.Scan(&r.ClusterName, &r.Namespace, &r.Kind, &r.Resource, &r.Age, &r.EventTime, &r.ExpiryDate); err != nil {
			return nil, fmt.Errorf("error scanning row: %v", err)
		}
		resources = append(resources, &r)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %v", err)
	}

	return resources, nil
}

// TrivyImages is the resolver for the trivyImages field.
func (r *queryResolver) TrivyImages(ctx context.Context, clusterName string) ([]*model.TrivyImage, error) {
	if r.DB == nil {
		return nil, fmt.Errorf("database connection is not initialized")
	}

	if clusterName == "" {
		return nil, fmt.Errorf("clusterName cannot be empty")
	}

	query := `
	SELECT id, cluster_name, artifact_name, vul_id, vul_pkg_id, vul_pkg_name, vul_installed_version, vul_fixed_version, vul_title, vul_severity, vul_published_date, vul_last_modified_date, ExpiryDate
	FROM trivyimage
	WHERE cluster_name = ?
	`

	rows, err := r.DB.QueryContext(ctx, query, clusterName)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %v", err)
	}
	defer rows.Close()

	var images []*model.TrivyImage
	for rows.Next() {
		var img model.TrivyImage
		if err := rows.Scan(&img.ID, &img.ClusterName, &img.ArtifactName, &img.VulID, &img.VulPkgID, &img.VulPkgName, &img.VulInstalledVersion, &img.VulFixedVersion, &img.VulTitle, &img.VulSeverity, &img.VulPublishedDate, &img.VulLastModifiedDate, &img.ExpiryDate); err != nil {
			return nil, fmt.Errorf("error scanning row: %v", err)
		}
		images = append(images, &img)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %v", err)
	}

	return images, nil
}

// DeprecatedAPIs is the resolver for the deprecatedAPIs field.
func (r *queryResolver) DeprecatedAPIs(ctx context.Context, clusterName string) ([]*model.DeprecatedAPI, error) {
	if r.DB == nil {
		return nil, fmt.Errorf("database connection is not initialized")
	}

	if clusterName == "" {
		return nil, fmt.Errorf("ClusterName cannot be empty")
	}

	query := `
	SELECT ClusterName, ObjectName, Description, Kind, Deprecated, Scope, EventTime, ExpiryDate
	FROM DeprecatedAPIs
	WHERE ClusterName = ?
	`

	rows, err := r.DB.QueryContext(ctx, query, clusterName)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %v", err)
	}
	defer rows.Close()
	var apis []*model.DeprecatedAPI
	for rows.Next() {
		var api model.DeprecatedAPI
		var deprecated uint8
		if err := rows.Scan(&api.ClusterName, &api.ObjectName, &api.Description, &api.Kind, &deprecated, &api.Scope, &api.EventTime, &api.ExpiryDate); err != nil {
			return nil, fmt.Errorf("error scanning row: %v", err)
		}
		deprecatedBool := deprecated != 0
		api.Deprecated = &deprecatedBool
		apis = append(apis, &api)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %v", err)
	}

	return apis, nil
}

// DeletedAPIs is the resolver for the deletedAPIs field.
func (r *queryResolver) DeletedAPIs(ctx context.Context, clusterName string) ([]*model.DeletedAPI, error) {
	if r.DB == nil {
		return nil, fmt.Errorf("database connection is not initialized")
	}

	if clusterName == "" {
		return nil, fmt.Errorf("ClusterName cannot be empty")
	}

	query := `
	SELECT ClusterName, ObjectName, Group, Kind, Version, Name, Deleted, Scope, EventTime, ExpiryDate
	FROM DeletedAPIs
	WHERE ClusterName = ?
	`
	rows, err := r.DB.QueryContext(ctx, query, clusterName)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %v", err)
	}
	defer rows.Close()
	var apis []*model.DeletedAPI
	for rows.Next() {
		var api model.DeletedAPI
		var deleted uint8
		if err := rows.Scan(&api.ClusterName, &api.ObjectName, &api.Group, &api.Kind, &api.Version, &api.Name, &deleted, &api.Scope, &api.EventTime, &api.ExpiryDate); err != nil {
			return nil, fmt.Errorf("error scanning row: %v", err)
		}
		deletedBool := deleted != 0
		api.Deleted = &deletedBool
		apis = append(apis, &api)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %v", err)
	}

	return apis, nil
}

// TrivySBOMs is the resolver for the trivySBOMs field.
func (r *queryResolver) TrivySBOMs(ctx context.Context, clusterName string) ([]*model.TrivySbom, error) {
	if r.DB == nil {
		return nil, fmt.Errorf("database connection is not initialized")
	}

	if clusterName == "" {
		return nil, fmt.Errorf("clusterName cannot be empty")
	}

	query := `
	SELECT id, cluster_name, image_name, package_name, package_url, bom_ref, serial_number, version, bom_format, ExpiryDate
	FROM trivysbom
	WHERE cluster_name = ?
	`

	rows, err := r.DB.QueryContext(ctx, query, clusterName)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %v", err)
	}
	defer rows.Close()

	var sboms []*model.TrivySbom
	for rows.Next() {
		var sbom model.TrivySbom
		if err := rows.Scan(&sbom.ID, &sbom.ClusterName, &sbom.ImageName, &sbom.PackageName, &sbom.PackageURL, &sbom.BomRef, &sbom.SerialNumber, &sbom.Version, &sbom.BomFormat, &sbom.ExpiryDate); err != nil {
			return nil, fmt.Errorf("error scanning row: %v", err)
		}
		sboms = append(sboms, &sbom)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %v", err)
	}

	return sboms, nil
}

// TrivyVulCount is the resolver for the trivyVulCount field.
func (r *queryResolver) TrivyVulCount(ctx context.Context, clusterName string, namespace string) (*model.ClusterNamespaceVulCount, error) {
	if r.DB == nil {
		return nil, fmt.Errorf("database connection is not initialized")
	}

	if clusterName == "" || namespace == "" {
		return nil, fmt.Errorf("clusterName and namespace cannot be empty")
	}

	query := `SELECT COUNT(*) FROM trivy_vul WHERE cluster_name = ? AND namespace = ?`

	var count int
	err := r.DB.QueryRowContext(ctx, query, clusterName, namespace).Scan(&count)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %v", err)
	}

	return &model.ClusterNamespaceVulCount{
		ClusterName: clusterName,
		Namespace:   namespace,
		VulCount:    count,
	}, nil
}

// TrivyMisconfigCount is the resolver for the trivyMisconfigCount field.
func (r *queryResolver) TrivyMisconfigCount(ctx context.Context, clusterName string, namespace string) (*model.ClusterNamespaceMisconfigCount, error) {
	if r.DB == nil {
		return nil, fmt.Errorf("database connection is not initialized")
	}

	if clusterName == "" || namespace == "" {
		return nil, fmt.Errorf("clusterName and namespace cannot be empty")
	}

	query := `SELECT COUNT(*) FROM trivy_misconfig WHERE cluster_name = ? AND namespace = ?`

	var count int
	err := r.DB.QueryRowContext(ctx, query, clusterName, namespace).Scan(&count)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %v", err)
	}

	return &model.ClusterNamespaceMisconfigCount{
		ClusterName:    clusterName,
		Namespace:      namespace,
		MisconfigCount: count,
	}, nil
}

// DeletedAPICount is the resolver for the deletedAPICount field.
func (r *queryResolver) DeletedAPICount(ctx context.Context, clusterName string) (*model.ClusterDeletedAPICount, error) {
	if r.DB == nil {
		return nil, fmt.Errorf("database connection is not initialized")
	}

	if clusterName == "" {
		return nil, fmt.Errorf("ClusterName cannot be empty")
	}

	query := `SELECT COUNT(*) FROM DeletedAPIs WHERE ClusterName = ?`

	var count int
	err := r.DB.QueryRowContext(ctx, query, clusterName).Scan(&count)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %v", err)
	}

	return &model.ClusterDeletedAPICount{
		ClusterName:     clusterName,
		DeletedAPICount: count,
	}, nil
}

// TrivyImageVulCount is the resolver for the trivyImageVulCount field.
func (r *queryResolver) TrivyImageCount(ctx context.Context, clusterName string) (*model.TrivyImageCount, error) {
	if r.DB == nil {
		return nil, fmt.Errorf("database connection is not initialized")
	}

	if clusterName == "" {
		return nil, fmt.Errorf("clusterName cannot be empty")
	}

	query := `SELECT COUNT(*) FROM trivyimage WHERE cluster_name = ?`

	var count int
	err := r.DB.QueryRowContext(ctx, query, clusterName).Scan(&count)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %v", err)
	}

	return &model.TrivyImageCount{
		ClusterName: clusterName,
		ImageCount:  count,
	}, nil
}

// DeprecatedAPICount is the resolver for the deprecatedAPICount field.
func (r *queryResolver) DeprecatedAPICount(ctx context.Context, clusterName string) (*model.ClusterDeprecatedAPICount, error) {
	if r.DB == nil {
		return nil, fmt.Errorf("database connection is not initialized")
	}

	if clusterName == "" {
		return nil, fmt.Errorf("ClusterName cannot be empty")
	}

	query := `SELECT COUNT(*) FROM DeprecatedAPIs WHERE ClusterName = ?`

	var count int
	err := r.DB.QueryRowContext(ctx, query, clusterName).Scan(&count)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %v", err)
	}

	return &model.ClusterDeprecatedAPICount{
		ClusterName:        clusterName,
		DeprecatedAPICount: count,
	}, nil
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
