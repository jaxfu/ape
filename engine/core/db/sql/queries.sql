-- name: GetComponentByComponentId :one
SELECT content
FROM components
WHERE component_id = ?;

-- name: InsertComponentMetadata :exec
INSERT INTO components(component_id, content)
VALUES (?, ?);