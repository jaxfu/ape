package internal

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/jaxfu/ape/components"
	"github.com/jaxfu/ape/engine/core/store/repo/db/sql/generated"
)

func (rp *Repository) getOpts(id components.ComponentId, t components.PropType) (components.PropTypeOptions, error) {
	switch t {
	case components.PROP_TYPE_INT:
		dbOpt, err := rp.QueryClient.GetOptsIntByPropId(context.Background(), strings.ToLower(id))
		if err != nil {
			return o, err
		}

		arr := new(bool)
		if dbOpt.Array.Valid {
			*arr = dbOpt.Array.Bool
		} else {
			arr = nil
		}

		size := new(uint)
		if dbOpt.Size.Valid {
			*size = uint(dbOpt.Size.Int64)
		} else {
			size = nil
		}

		min := new(int)
		if dbOpt.Min.Valid {
			*min = int(dbOpt.Min.Int64)
		} else {
			min = nil
		}

		max := new(int)
		if dbOpt.Max.Valid {
			*max = int(dbOpt.Max.Int64)
		} else {
			max = nil
		}

		o = components.Opts{
			Int: &components.OptsFieldsInt{
				Array: arr,
				Size:  size,
				Min:   min,
				Max:   max,
			},
		}
	case components.PROP_TYPE_UINT:
		dbOpt, err := rp.QueryClient.GetOptsUintByPropId(context.Background(), strings.ToLower(id))
		if err != nil {
			return o, err
		}

		arr := new(bool)
		if dbOpt.Array.Valid {
			*arr = dbOpt.Array.Bool
		} else {
			arr = nil
		}

		size := new(uint)
		if dbOpt.Size.Valid {
			*size = uint(dbOpt.Size.Int64)
		} else {
			size = nil
		}

		min := new(uint)
		if dbOpt.Min.Valid {
			*min = uint(dbOpt.Min.Int64)
		} else {
			min = nil
		}

		max := new(uint)
		if dbOpt.Max.Valid {
			*max = uint(dbOpt.Max.Int64)
		} else {
			max = nil
		}

		o = components.Opts{
			Uint: &components.OptsFieldsUint{
				Array: arr,
				Size:  size,
				Min:   min,
				Max:   max,
			},
		}
	case components.PROP_TYPE_FLOAT:
		dbOpt, err := rp.QueryClient.GetOptsFloatByPropId(context.Background(), strings.ToLower(id))
		if err != nil {
			return o, err
		}

		arr := new(bool)
		if dbOpt.Array.Valid {
			*arr = dbOpt.Array.Bool
		} else {
			arr = nil
		}

		prc := new(string)
		if dbOpt.Precision.String != "" {
			*prc = dbOpt.Precision.String
		} else {
			prc = nil
		}

		mn := new(float64)
		if dbOpt.Min.Valid {
			*mn = dbOpt.Min.Float64
		} else {
			mn = nil
		}

		mx := new(float64)
		if dbOpt.Max.Valid {
			*mx = dbOpt.Max.Float64
		} else {
			mx = nil
		}

		o = components.Opts{
			Float: &components.OptsFieldsFloat{
				Array:     arr,
				Precision: prc,
				Min:       mn,
				Max:       mx,
			},
		}
	case components.PROP_TYPE_TEXT:
		dbOpt, err := rp.QueryClient.GetOptsTextByPropId(context.Background(), strings.ToLower(id))
		if err != nil {
			return o, err
		}

		arr := new(bool)
		if dbOpt.Array.Valid {
			*arr = dbOpt.Array.Bool
		} else {
			arr = nil
		}

		rgx := new(string)
		if dbOpt.Regex.String != "" {
			*rgx = dbOpt.Regex.String
		} else {
			rgx = nil
		}

		mn := new(uint)
		if dbOpt.MinLength.Valid {
			*mn = uint(dbOpt.MinLength.Int64)
		} else {
			mn = nil
		}

		mx := new(uint)
		if dbOpt.MaxLength.Valid {
			*mx = uint(dbOpt.MaxLength.Int64)
		} else {
			mx = nil
		}

		aln := new(bool)
		if dbOpt.Alnum.Valid {
			*aln = dbOpt.Alnum.Bool
		} else {
			aln = nil
		}

		alp := new(bool)
		if dbOpt.Alpha.Valid {
			*alp = dbOpt.Alpha.Bool
		} else {
			alp = nil
		}

		num := new(bool)
		if dbOpt.Num.Valid {
			*num = dbOpt.Num.Bool
		} else {
			num = nil
		}

		o = components.Opts{
			Text: &components.OptsFieldsText{
				Array:     arr,
				MinLength: mn,
				MaxLength: mx,
				Regex:     rgx,
				Alnum:     aln,
				Alpha:     alp,
				Num:       num,
			},
		}
	case components.PROP_TYPE_BLOB:
		dbOpt, err := rp.QueryClient.GetOptsBlobByPropId(context.Background(), strings.ToLower(id))
		if err != nil {
			return o, err
		}

		arr := new(bool)
		if dbOpt.Array.Valid {
			*arr = dbOpt.Array.Bool
		} else {
			arr = nil
		}

		mn := new(uint)
		if dbOpt.MinSize.Valid {
			*mn = uint(dbOpt.MinSize.Int64)
		} else {
			mn = nil
		}

		mx := new(uint)
		if dbOpt.MaxSize.Valid {
			*mx = uint(dbOpt.MaxSize.Int64)
		} else {
			mx = nil
		}

		o = components.Opts{
			Blob: &components.OptsFieldsBlob{
				Array:   arr,
				MinSize: mn,
				MaxSize: mx,
			},
		}
	case components.PROP_TYPE_REF:
		dbOpt, err := rp.QueryClient.GetRefByComponentId(context.Background(), strings.ToLower(id))
		if err != nil {
			return o, err
		}

		o = components.Opts{
			Ref: &components.OptsFieldsRef{
				TargetId: &dbOpt.RawTargetID,
			},
		}
	}

	return o, nil
}

func (rp *Repository) storeOpts(prop components.Prop, parentId components.ComponentId) error {
	switch prop.PropType {
	case components.PROP_TYPE_INT:
		if prop.Opts.Int != nil {
			opts := prop.Opts.Int

			arr := sql.NullBool{}
			if opts.Array != nil {
				arr.Bool = *opts.Array
				arr.Valid = true
			}

			sz := sql.NullInt64{}
			if opts.Size != nil {
				sz.Int64 = int64(*opts.Size)
				sz.Valid = true
			}

			mn := sql.NullInt64{}
			if opts.Min != nil {
				mn.Int64 = int64(*opts.Min)
				mn.Valid = true
			}

			mx := sql.NullInt64{}
			if opts.Max != nil {
				mx.Int64 = int64(*opts.Max)
				mx.Valid = true
			}

			prms := generated.InsertOptsIntParams{
				ParentID: parentId,
				Array:    arr,
				Size:     sz,
				Min:      mn,
				Max:      mx,
			}

			if err := rp.QueryClient.InsertOptsInt(
				context.Background(),
				prms,
			); err != nil {
				return fmt.Errorf("InsertOptsInt: %+v", err)
			}
		}
	case components.PROP_TYPE_UINT:
		if prop.Opts.Uint != nil {
			opts := prop.Opts.Uint

			arr := sql.NullBool{}
			if opts.Array != nil {
				arr.Bool = *opts.Array
				arr.Valid = true
			}

			sz := sql.NullInt64{}
			if opts.Size != nil {
				sz.Int64 = int64(*opts.Size)
				sz.Valid = true
			}

			mn := sql.NullInt64{}
			if opts.Min != nil {
				mn.Int64 = int64(*opts.Min)
				mn.Valid = true
			}

			mx := sql.NullInt64{}
			if opts.Max != nil {
				mx.Int64 = int64(*opts.Max)
				mx.Valid = true
			}

			prms := generated.InsertOptsUintParams{
				ParentID: parentId,
				Array:    arr,
				Size:     sz,
				Min:      mn,
				Max:      mx,
			}

			if err := rp.QueryClient.InsertOptsUint(
				context.Background(),
				prms,
			); err != nil {
				return fmt.Errorf("InsertOptsUint: %+v", err)
			}
		}
	case components.PROP_TYPE_FLOAT:
		if prop.Opts.Float != nil {
			opts := prop.Opts.Float

			arr := sql.NullBool{}
			if opts.Array != nil {
				arr.Bool = *opts.Array
				arr.Valid = true
			}

			prc := sql.NullString{}
			if opts.Precision != nil {
				prc.String = *opts.Precision
				prc.Valid = true
			}

			mn := sql.NullFloat64{}
			if opts.Min != nil {
				mn.Float64 = *opts.Min
				mn.Valid = true
			}

			mx := sql.NullFloat64{}
			if opts.Max != nil {
				mx.Float64 = *opts.Max
				mx.Valid = true
			}

			prms := generated.InsertOptsFloatParams{
				ParentID:  parentId,
				Array:     arr,
				Precision: prc,
				Min:       mn,
				Max:       mx,
			}

			if err := rp.QueryClient.InsertOptsFloat(
				context.Background(),
				prms,
			); err != nil {
				return fmt.Errorf("InsertOptsFloat: %+v", err)
			}
		}
	case components.PROP_TYPE_TEXT:
		if prop.Opts.Text != nil {
			opts := prop.Opts.Text

			arr := sql.NullBool{}
			if opts.Array != nil {
				arr.Bool = *opts.Array
				arr.Valid = true
			}

			rgx := sql.NullString{}
			if opts.Regex != nil {
				rgx.String = *opts.Regex
				rgx.Valid = true
			}

			mn := sql.NullInt64{}
			if opts.MinLength != nil {
				mn.Int64 = int64(*opts.MinLength)
				mn.Valid = true
			}

			mx := sql.NullInt64{}
			if opts.MaxLength != nil {
				mx.Int64 = int64(*opts.MaxLength)
				mx.Valid = true
			}

			aln := sql.NullBool{}
			if opts.Alnum != nil {
				aln.Bool = *opts.Alnum
				aln.Valid = true
			}

			alp := sql.NullBool{}
			if opts.Alpha != nil {
				alp.Bool = *opts.Alpha
				alp.Valid = true
			}

			num := sql.NullBool{}
			if opts.Num != nil {
				num.Bool = *opts.Num
				num.Valid = true
			}

			prms := generated.InsertOptsTextParams{
				ParentID:  parentId,
				Array:     arr,
				Regex:     rgx,
				MinLength: mn,
				MaxLength: mx,
				Alnum:     aln,
				Alpha:     alp,
				Num:       num,
			}

			if err := rp.QueryClient.InsertOptsText(
				context.Background(),
				prms,
			); err != nil {
				return fmt.Errorf("InsertOptsText: %+v", err)
			}
		}
	case components.PROP_TYPE_BLOB:
		if prop.Opts.Blob != nil {
			opts := prop.Opts.Blob

			arr := sql.NullBool{}
			if opts.Array != nil {
				arr.Bool = *opts.Array
				arr.Valid = true
			}

			mn := sql.NullInt64{}
			if opts.MinSize != nil {
				mn.Int64 = int64(*opts.MinSize)
				mn.Valid = true
			}

			mx := sql.NullInt64{}
			if opts.MaxSize != nil {
				mx.Int64 = int64(*opts.MaxSize)
				mx.Valid = true
			}

			prms := generated.InsertOptsBlobParams{
				ParentID: parentId,
				Array:    arr,
				MinSize:  mn,
				MaxSize:  mx,
			}

			if err := rp.QueryClient.InsertOptsBlob(
				context.Background(),
				prms,
			); err != nil {
				return fmt.Errorf("InsertOptsFloat: %+v", err)
			}
		}
	case components.PROP_TYPE_REF:
		if prop.Opts.Ref != nil {
			opts := prop.Opts.Ref

			targetId := ""
			if opts.TargetId != nil && strings.TrimSpace(*opts.TargetId) != "" {
				targetId = *opts.TargetId
			} else {
				return fmt.Errorf("no targetId given for Prop Ref '%s'", prop.Name)
			}

			// validate ref before inserting
			linked := sql.NullString{}
			if _, err := rp.QueryClient.GetComponentByComponentId(
				context.Background(),
				strings.ToLower(targetId),
			); err != nil {
				if !errors.Is(err, sql.ErrNoRows) {
					fmt.Printf("error searching for Ref target '%s': %+v\n", strings.ToLower(targetId), err)
				}
			} else {
				linked.String = strings.ToLower(targetId)
				linked.Valid = true
			}

			prms := generated.InsertRefParams{
				ComponentID:    parentId,
				RawTargetID:    targetId,
				LinkedTargetID: linked,
			}

			if err := rp.QueryClient.InsertRef(
				context.Background(),
				prms,
			); err != nil {
				return fmt.Errorf("Repo.InsertRef: %+v", err)
			}
		}

	}

	return nil
}
