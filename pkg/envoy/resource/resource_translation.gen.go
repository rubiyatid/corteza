package resource

// This file is auto-generated.
//
// Changes to this file may cause incorrect behavior and will be lost if
// the code is regenerated.
//

// Definitions file that controls how this file is generated:
// - compose.module-field.yaml
// - compose.module.yaml
// - compose.namespace.yaml
// - compose.page.yaml

import (
	systemTypes "github.com/cortezaproject/corteza-server/system/types"
)

func (r *ComposeModuleField) EncodeTranslations() ([]*ResourceTranslation, error) {
	out := make([]*ResourceTranslation, 0, 10)

	rr := r.Res.EncodeTranslations()
	rr.SetLanguage(defaultLanguage)
	res, ref, pp := r.ResourceTranslationParts()
	out = append(out, NewResourceTranslation(systemTypes.FromLocale(rr), res, ref, pp...))

	return out, nil
}

func (r *ComposeModule) EncodeTranslations() ([]*ResourceTranslation, error) {
	out := make([]*ResourceTranslation, 0, 10)

	rr := r.Res.EncodeTranslations()
	rr.SetLanguage(defaultLanguage)
	res, ref, pp := r.ResourceTranslationParts()
	out = append(out, NewResourceTranslation(systemTypes.FromLocale(rr), res, ref, pp...))

	tmp, err := r.encodeTranslations()
	return append(out, tmp...), err

}

func (r *ComposeNamespace) EncodeTranslations() ([]*ResourceTranslation, error) {
	out := make([]*ResourceTranslation, 0, 10)

	rr := r.Res.EncodeTranslations()
	rr.SetLanguage(defaultLanguage)
	res, ref, pp := r.ResourceTranslationParts()
	out = append(out, NewResourceTranslation(systemTypes.FromLocale(rr), res, ref, pp...))

	return out, nil
}

func (r *ComposePage) EncodeTranslations() ([]*ResourceTranslation, error) {
	out := make([]*ResourceTranslation, 0, 10)

	rr := r.Res.EncodeTranslations()
	rr.SetLanguage(defaultLanguage)
	res, ref, pp := r.ResourceTranslationParts()
	out = append(out, NewResourceTranslation(systemTypes.FromLocale(rr), res, ref, pp...))

	tmp, err := r.encodeTranslations()
	return append(out, tmp...), err

}
