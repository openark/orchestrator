package rqlite

import (
	"fmt"

	"github.com/mkideal/cli"
	"github.com/mkideal/pkg/textutil"
)

// Rows represents query result
type Rows struct {
	Columns []string        `json:"columns"`
	Types   []string        `json:"types"`
	Values  [][]interface{} `json:"values"`
	Time    float64         `json:"time"`
	Error   string          `json:"error,omitempty"`
}

// RowCount implements textutil.Table interface
func (r *Rows) RowCount() int {
	return len(r.Values) + 1
}

// ColCount implements textutil.Table interface
func (r *Rows) ColCount() int {
	return len(r.Columns)
}

// Get implements textutil.Table interface
func (r *Rows) Get(i, j int) string {
	if i == 0 {
		if j >= len(r.Columns) {
			return ""
		}
		return r.Columns[j]
	}

	if r.Values == nil {
		return "NULL"
	}

	if i-1 >= len(r.Values) {
		return "NULL"
	}
	if j >= len(r.Values[i-1]) {
		return "NULL"
	}
	return fmt.Sprintf("%v", r.Values[i-1][j])
}

func (r *Rows) validate() error {
	if r.Error != "" {
		return fmt.Errorf(r.Error)
	}
	if r.Columns == nil || r.Types == nil {
		return fmt.Errorf("unexpected result")
	}
	return nil
}

// headerRenderStyle render the header of result
type headerRenderStyle struct {
	textutil.DefaultStyle
}

func (render headerRenderStyle) CellRender(row, col int, cell string, cw *textutil.ColorWriter) {
	if row != 0 {
		fmt.Fprintf(cw, cell)
	} else {
		fmt.Fprintf(cw, cw.Color.Cyan(cell))
	}
}

var headerRender = &headerRenderStyle{}

type queryResponse struct {
	Results []*Rows `json:"results"`
	Error   string  `json:"error,omitempty"`
	Time    float64 `json:"time"`
}

func query(ctx *cli.Context, cmd, line string, argv *argT) error {
	urlStr := fmt.Sprintf("%s://%s:%d/db/query?timings", argv.Protocol, argv.Host, argv.Port)
	fmt.Println(urlStr)
	ret := &queryResponse{}
	if err := sendRequest(ctx, urlStr, line, ret); err != nil {
		return err
	}
	if ret.Error != "" {
		return fmt.Errorf(ret.Error)
	}
	if len(ret.Results) != 1 {
		// NOTE:What's happen? ret.Results.length MUST be 1
		return fmt.Errorf("unexpected results length: %d", len(ret.Results))
	}

	result := ret.Results[0]
	if err := result.validate(); err != nil {
		return err
	}
	textutil.WriteTable(ctx, result, headerRender)
	return nil
}
